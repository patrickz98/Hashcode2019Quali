package pizza

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../simple"
)

type parameters struct {
	Rows        int
	Columns     int
	Ingredients int
	MaxCells    int
}

type Coordinate struct {
	Row    int
	Column int
}

type Cell struct {
	Slice *Slice
	Type rune
}

type Pizza struct {
	Ingredients int
	MaxCells    int
	Cells       map[Coordinate] *Cell
	Row         Vector
	Column      Vector
}

func (pizza Pizza) PrintParams() {

	fmt.Printf("Ingredients: %d\n", pizza.Ingredients)
	fmt.Printf("MaxCells: %d\n", pizza.MaxCells)
	fmt.Printf("Rows: %d\n", pizza.Row.Length())
	fmt.Printf("Columns: %d\n", pizza.Column.Length())
}

func (pizza Pizza) Size() int {

	return pizza.Column.Length() * pizza.Row.Length()
}

func (pizza Pizza) Traversal() []Coordinate {

	coordinates := make([]Coordinate, pizza.Size())

	for iny, row := range pizza.Row.Range() {
		for inx, col := range pizza.Column.Range() {
			index := (iny * pizza.Column.Length()) + inx
			coordinates[ index ] = Coordinate{Row: row, Column: col}
		}
	}

	return coordinates
}

func (pizza Pizza) PrintPizza() {

	for _, xy := range pizza.Traversal() {

		cell := pizza.Cells[ xy ]
		fmt.Print(string(cell.Type))
	}

	fmt.Println()
}

func (pizza Pizza) SlicesAsString() string {
	width := pizza.Column.Length() * 2 + 1
	height := pizza.Row.Length() * 2 + 1

	field := make([][]rune, height)

	for iny := range field {
		field[ iny ] = make([]rune, width)
		for inx := range field[ iny ] {
			field[ iny ][ inx ] = ' '
		}
	}

	slices := make([]Slice, 0)

	for iny, yy := range pizza.Row.Range() {
		for inx, xx := range pizza.Column.Range() {
			coord := Coordinate{Row: yy, Column: xx}
			cell := pizza.Cells[ coord ]

			field[ iny * 2 + 1 ][ inx * 2 + 1 ] = cell.Type

			if cell.Slice != nil {
				slices = append(slices, *cell.Slice)
			}
		}
	}

	for _, sli := range slices {

		t := (sli.Row.Start    - pizza.Row.Start) * 2 + 1
		b := (sli.Row.End      - pizza.Row.Start) * 2 + 1
		l := (sli.Column.Start - pizza.Column.Start) * 2
		r := (sli.Column.End   - pizza.Column.Start) * 2 + 1

		// piz.PrintSlicesPlain()
		// sli.PrintVector()

		// fmt.Printf("t=%d b=%d\n", t, b)
		// fmt.Printf("l=%d r=%d\n", l, r)

		// break

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

	text := ""

	for iny := range field {
		text += string(field[ iny ]) + "\n"
	}

	return text
}

func (pizza Pizza) PrintSlices() {

	fmt.Print(pizza.SlicesAsString())
}

func (pizza Pizza) PrintSlicesToFile(path string) {

	text := pizza.SlicesAsString()
	bytes := []byte(text)

	err := ioutil.WriteFile(path, bytes, 0644)
	simple.CheckErr(err)
}

func (pizza Pizza) Score() (total int, covered int, score float32) {

	total = pizza.Size()

	// TODO: make nice.
	tmp := make(map[string] int)

	for _, xy := range pizza.Traversal() {
		cell := pizza.Cells[ xy ]

		if cell.Slice != nil {
			col := cell.Slice.Row.Stringify()
			row := cell.Slice.Column.Stringify()

			tmp[ col + row ] = cell.Slice.Size()
		}
	}

	covered = 0

	for _, inx := range tmp {
		covered += inx
	}

	score = float32(covered) / float32(total)

	return total, covered, score
}

func (pizza Pizza) PrintScore() {

	total, count, score := pizza.Score()

	fmt.Printf("total cells: %d\n", total)
	fmt.Printf("covered: %d\n", count)
	fmt.Printf("percent: %.2f%%\n", score * 100)
}

func (pizza Pizza) PrintVector(row Vector, column Vector) {

	for _, iny := range row.Range() {
		for _, inx := range column.Range() {

			xy := Coordinate{Row: iny, Column: inx}
			cell := pizza.Cells[ xy ]
			fmt.Print(string(cell.Type))
		}

		fmt.Println()
	}
}

func (pizza Pizza) PrintVectors() {

	fmt.Printf("vectorR := pizza.Vector{Start: %d, End: %d}\n", pizza.Row.Start, pizza.Row.End)
	fmt.Printf("vectorC := pizza.Vector{Start: %d, End: %d}\n", pizza.Column.Start, pizza.Column.End)
}

func (pizza Pizza) PrintSlicesPlain() {

	for _, xy := range pizza.Traversal() {
		cell := pizza.Cells[ xy ]

		if cell.Slice != nil {
			cell.Slice.PrintVector()
		}
	}
}

func (pizza *Pizza) AddSlice(slice *Slice) {

	for _, xy := range slice.Traversal() {
		cell := pizza.Cells[ xy ]
		cell.Slice = slice
	}
}

func (pizza *Pizza) RemoveSlice(slice *Slice) {

	for _, xy := range slice.Traversal() {
		cell := pizza.Cells[ xy ]
		cell.Slice = nil
	}
}

func initParams(head string) parameters {

	line := strings.TrimSuffix(head, "\n")
	parts := strings.Split(line, " ")

	paramsArray := make([]int64, 4)

	for inx, str := range parts {
		val, _ := strconv.ParseInt(str, 10, 64)
		paramsArray[ inx ] = val
	}

	params := parameters{
		Rows:        int(paramsArray[ 0 ]),
		Columns:     int(paramsArray[ 1 ]),
		Ingredients: int(paramsArray[ 2 ]),
		MaxCells:    int(paramsArray[ 3 ]),
	}

	return params
}

func initPizza(params parameters, lines []string) Pizza {

	cells := make(map[Coordinate] *Cell)

	for iny, line := range lines {
		line = strings.TrimSuffix(line, "\n")
		runes := []rune(line)

		for inx, val := range runes {

			coordinate := Coordinate{Row: iny, Column: inx}
			cells[ coordinate ] = &Cell{Type: val}
		}
	}

	return Pizza{
		Ingredients: params.Ingredients,
		MaxCells:    params.MaxCells,
		Cells:       cells,
		Row:         Vector{Start: 0, End: params.Rows - 1},
		Column:      Vector{Start: 0, End: params.Columns - 1},
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
