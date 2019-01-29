package pizza

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../simple"
)

type parameters struct {
	Rows       int
	Columns    int
	Ingredient int
	MaxCells   int
}

type Cell struct {
	Slice *Slice
	Type rune
}

type Pizza struct {
	*parameters
	Cells  [][]Cell
	Slices []*Slice
}

func (piz Pizza) PrintParams() {

	bytes, err := simple.PrettyJson(piz.parameters)
	simple.CheckErr(err)

	fmt.Println(string(bytes))
}

func (piz Pizza) PrintPizza() {

	row := Vector{Start: 0, End: piz.Rows - 1}
	columns := Vector{Start: 0, End: piz.Columns - 1}

	piz.PrintVector(row, columns)
}

// func (piz Pizza) PrintPizzaCells() {
//
// 	for iny := piz.Rows.Start; iny < piz.Rows.End+1; iny++ {
// 		columns := piz.Cells[ iny ][ piz.Columns.Start : piz.Columns.End+1 ]
//
// 		for _, cell := range columns {
// 			fmt.Printf("%c %t ", cell.Type, cell.Slice != nil)
// 		}
//
// 		fmt.Println()
// 	}
// }

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

func initParams(head string) parameters {

	line := strings.TrimSuffix(head, "\n")
	parts := strings.Split(line, " ")

	paramsArray := make([]int64, 4)

	for inx, str := range parts {
		val, _ := strconv.ParseInt(str, 10, 64)
		paramsArray[ inx ] = val
	}

	params := parameters{
		Rows:       int(paramsArray[ 0 ]),
		Columns:    int(paramsArray[ 1 ]),
		Ingredient: int(paramsArray[ 2 ]),
		MaxCells:   int(paramsArray[ 3 ]),
	}

	return params
}

func initPizza(params parameters, lines []string) Pizza {

	data := make([][]Cell, params.Rows)

	for inx := range data {
		data[ inx ] = make([]Cell, params.Columns)
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
		Slices:     []*Slice{},
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
