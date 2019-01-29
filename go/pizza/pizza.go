package pizza

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../simple"
)

type parameters struct {
	Rows       Vector
	Columns    Vector
	Ingredient int
	MaxCells   int
}

type Slice struct {
	Row    Vector
	Column Vector
}

func (slice Slice) Size() int {

	return slice.Row.Size(slice.Column)
}

func (slice Slice) PrintVector() {
	fmt.Printf("row=%s column=%s\n", slice.Row.Stringify(), slice.Column.Stringify())
}

type Cell struct {
	Slice *Slice
	Type rune
}

type Pizza struct {
	*parameters
	Cells  [][]Cell
	Slices []Slice
}

func (piz Pizza) PrintParams() {

	bytes, err := simple.PrettyJson(piz.parameters)
	simple.CheckErr(err)

	fmt.Println(string(bytes))
}

func (piz Pizza) PrintPizza() {

	for iny := piz.Rows.Start; iny < piz.Rows.End+1; iny++ {
		columns := piz.Cells[ iny ][ piz.Columns.Start : piz.Columns.End+1 ]

		for _, cell := range columns {
			fmt.Print(string(cell.Type))
		}

		fmt.Println()
	}
}

func (piz Pizza) PrintPizzaCells() {

	for iny := piz.Rows.Start; iny < piz.Rows.End+1; iny++ {
		columns := piz.Cells[ iny ][ piz.Columns.Start : piz.Columns.End+1 ]

		for _, cell := range columns {
			fmt.Printf("%c %t ", cell.Type, cell.Slice != nil)
		}

		fmt.Println()
	}
}

func (piz Pizza) PrintSlices() {

	width := piz.Columns.Length() * 2 + 1
	height := piz.Rows.Length() * 2 + 1

	field := make([][]rune, height)

	for inx := range field {
		field[ inx ] = make([]rune, width)
	}

	for iny := range field {
		for inx := range field[ iny ] {
			field[ iny ][ inx ] = ' '
		}
	}

	for iny := range piz.Rows.Range() {
		for inx := range piz.Columns.Range() {
			field[ iny * 2 + 1 ][ inx * 2 +1 ] = piz.Cells[ iny ][ inx ].Type
		}
	}

	for _, sli := range piz.Slices {

		t := sli.Row.Start * 2 + 1
		b := sli.Row.End * 2 + 1

		l := sli.Column.Start * 2
		r := sli.Column.End * 2 + 1

		horizontalLenth := sli.Column.Length() * 2

		for iny := t; iny < b + 1; iny = iny + 2 {
			field[ iny ][ l ] = '|'
			field[ iny ][ l + horizontalLenth ] = '|'
		}

		for inx := l + 1; inx < r + 1; inx = inx + 2 {
			field[ t - 1 ][ inx ] = '-'
			field[ b + 1 ][ inx ] = '-'
		}
	}

	for iny := range field {
		fmt.Println(string(field[ iny ]))
	}
}

func (piz Pizza) PrintScore() {

	count := 0

	for _, sli := range piz.Slices {
		count += sli.Size()
	}

	total := piz.Rows.Size(piz.Columns)
	fmt.Printf("total cells: %d\n", total)
	fmt.Printf("covered: %d\n", count)
	fmt.Printf("percent: %.2f%%\n", (float32(count) / float32(total)) * 100)
}

func (piz Pizza) CutPossible() bool {

	return piz.Rows.CutPossible() || piz.Columns.CutPossible()
}

func (piz Pizza) CutPeace(row Vector, column Vector) *Pizza {

	pizza := &Pizza{
		parameters: &parameters{
			Rows:       row,
			Columns:    column,
			Ingredient: piz.Ingredient,
			MaxCells:   piz.MaxCells,
		},
		Cells:  piz.Cells,
		Slices: []Slice{},
	}

	return pizza
}

func (piz Pizza) Cut() []*Pizza {

	parts := make([]*Pizza, 4)

	if piz.Rows.CutPossible() && piz.Columns.CutPossible() {

		r1, r2 := piz.Rows.Cut()
		c1, c2 := piz.Columns.Cut()

		plt := piz.CutPeace(r1, c1)
		prt := piz.CutPeace(r1, c2)
		plb := piz.CutPeace(r2, c1)
		prb := piz.CutPeace(r2, c2)

		parts[ 0 ] = plt
		parts[ 1 ] = prt
		parts[ 2 ] = plb
		parts[ 3 ] = prb

		return parts
	}

	if piz.Rows.CutPossible() {

		r1, r2 := piz.Rows.Cut()

		pt := piz.CutPeace(r1, piz.Columns)
		pb := piz.CutPeace(r2, piz.Columns)

		parts[ 0 ] = pt
		parts[ 2 ] = pb

		return parts
	}

	if piz.Columns.CutPossible() {

		c1, c2 := piz.Columns.Cut()

		pl := piz.CutPeace(piz.Rows, c1)
		pr := piz.CutPeace(piz.Rows, c2)

		parts[ 0 ] = pl
		parts[ 1 ] = pr

		return parts
	}

	return parts
}

func (piz Pizza) PrintVector(row Vector, column Vector) {

	for iny := row.Start; iny < row.End+1; iny++ {
		columns := piz.Cells[ iny ][ column.Start : column.End+1 ]

		for _, cell := range columns {
			fmt.Print(string(cell.Type))
		}

		fmt.Println()
	}
}

func (piz Pizza) PrintSlice(slice Slice) {

	piz.PrintVector(slice.Row, slice.Column)
}

func (piz *Pizza) AddSlice(slice Slice) {

	// fmt.Println("Add Slice")

	row := slice.Row
	column := slice.Column

	for iny := row.Start; iny < row.End+1; iny++ {
		for inx := column.Start; inx < column.End+1; inx++ {

			piz.Cells[ iny ][ inx ].Slice = &slice
		}
	}

	piz.Slices = append(piz.Slices, slice)
}


func initParams(head string) parameters {

	line := strings.TrimSuffix(head, "\n")
	parts := strings.Split(line, " ")

	paramsArray := make([]int64, 4)

	for inx, str := range parts {
		val, _ := strconv.ParseInt(str, 10, 64)
		paramsArray[ inx ] = val
	}

	rVector := Vector{Start: 0, End: int(paramsArray[ 0 ]) - 1}
	cVector := Vector{Start: 0, End: int(paramsArray[ 1 ]) - 1}

	params := parameters{
		Rows:       rVector,
		Columns:    cVector,
		Ingredient: int(paramsArray[ 2 ]),
		MaxCells:   int(paramsArray[ 3 ]),
	}

	return params
}

func initPizza(params parameters, lines []string) Pizza {

	data := make([][]Cell, params.Rows.Length())

	for inx := range data {
		data[ inx ] = make([]Cell, params.Columns.Length())
	}

	for inx, line := range lines {
		line = strings.TrimSuffix(line, "\n")
		runes := []rune(line)

		for iny, val := range runes {
			data[ inx ][ iny ] = Cell{Type: val}
		}
	}

	return Pizza{
		parameters: &params,
		Cells:      data,
		Slices:     []Slice{},
	}
}

func NewPizza(path string) Pizza {
	dat, err := ioutil.ReadFile(path)
	simple.CheckErr(err)

	lines := strings.SplitAfter(string(dat), "\n")

	head, lines := lines[ 0 ], lines[ 1:]
	params := initParams(head)

	return initPizza(params, lines)
}
