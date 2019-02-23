package pizza

import (
	"../simple"
	"io/ioutil"
	"strconv"
	"strings"
)

type parameters struct {
	Rows        int
	Columns     int
	Ingredients int
	MaxCells    int
}

func initParams(head string) parameters {

	line := strings.TrimSuffix(head, "\n")
	parts := strings.Split(line, " ")

	paramsArray := make([]int64, 4)

	for inx, str := range parts {
		val, _ := strconv.ParseInt(str, 10, 64)
		paramsArray[inx] = val
	}

	params := parameters{
		Rows:        int(paramsArray[0]),
		Columns:     int(paramsArray[1]),
		Ingredients: int(paramsArray[2]),
		MaxCells:    int(paramsArray[3]),
	}

	return params
}

func initPizza(params parameters, lines []string) Pizza {

	cells := make(map[Coordinate]*Cell)

	for iny, line := range lines {
		line = strings.TrimSuffix(line, "\n")
		runes := []rune(line)

		for inx, val := range runes {

			coordinate := Coordinate{Row: iny, Column: inx}
			cells[coordinate] = &Cell{Type: val}
		}
	}

	return Pizza{
		Ingredients: params.Ingredients,
		MaxCells:    params.MaxCells,
		Cells:       cells,
		Row:         Vector{Start: 0, End: params.Rows - 1},
		Column:      Vector{Start: 0, End: params.Columns - 1},
		ScoreTotal:  0,
	}
}

func NewPizza(path string) Pizza {
	dat, err := ioutil.ReadFile(path)
	simple.CheckErr(err)

	lines := strings.SplitAfter(string(dat), "\n")

	head, lines := lines[0], lines[1:]
	params := initParams(head)

	return initPizza(params, lines)
}
