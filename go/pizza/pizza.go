package pizza

import (
	"../simple"
	"fmt"
	"io/ioutil"
)

type Coordinate struct {
	Row    int
	Column int
}

func (xy Coordinate) Stringify() string {
	return fmt.Sprintf("(%d, %d)", xy.Row, xy.Column)
}

func (xy Coordinate) AddTo(row int, column int) Coordinate {
	xy.Row += row
	xy.Column += column
	return xy
}

type Cell struct {
	Slice *Slice
	Type  rune
}

type Pizza struct {
	Ingredients int
	MaxCells    int
	Cells       map[Coordinate]*Cell
	Row         Vector
	Column      Vector

	ScoreTotal int
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
			coordinates[index] = Coordinate{Row: row, Column: col}
		}
	}

	return coordinates
}

func (pizza Pizza) TraversalNotSlicedCells() []Coordinate {

	coordinates := make([]Coordinate, 0)

	for _, xy := range pizza.Traversal() {

		if pizza.Cells[xy].Slice == nil {
			coordinates = append(coordinates, xy)
		}
	}

	return coordinates
}

func (pizza Pizza) SlicesAsString(mark bool) string {
	width := pizza.Column.Length()*2 + 1
	height := pizza.Row.Length()*2 + 1

	field := make([][]rune, height)

	for iny := range field {
		field[iny] = make([]rune, width)
		for inx := range field[iny] {
			field[iny][inx] = ' '
		}
	}

	slices := make([]Slice, 0)

	for iny, yy := range pizza.Row.Range() {
		for inx, xx := range pizza.Column.Range() {
			coord := Coordinate{Row: yy, Column: xx}
			cell := pizza.Cells[coord]

			// if mark && cell.Slice == nil {
			// 	field[ iny * 2 + 1 ][ inx * 2 + 1 ] = '*'
			// } else {
			// 	field[ iny * 2 + 1 ][ inx * 2 + 1 ] = cell.Type
			// }

			if mark {
				if cell.Slice == nil {
					field[iny*2+1][inx*2+1] = cell.Type
				} else {
					field[iny*2+1][inx*2+1] = ' '
				}
			} else {
				field[iny*2+1][inx*2+1] = cell.Type
			}

			if cell.Slice != nil {
				slices = append(slices, *cell.Slice)
			}
		}
	}

	for _, sli := range slices {

		t := (sli.Row.Start-pizza.Row.Start)*2 + 1
		b := (sli.Row.End-pizza.Row.Start)*2 + 1
		l := (sli.Column.Start - pizza.Column.Start) * 2
		r := (sli.Column.End-pizza.Column.Start)*2 + 1

		horizontalLenth := sli.Column.Length() * 2

		for iny := t; iny < b+1; iny = iny + 2 {
			field[iny][l] = '|'
			field[iny][l+horizontalLenth] = '|'
		}

		for inx := l + 1; inx < r+1; inx = inx + 2 {
			field[t-1][inx] = '-'
			field[b+1][inx] = '-'
		}
	}

	text := ""

	for iny := range field {
		text += string(field[iny]) + "\n"
	}

	return text
}

func (pizza Pizza) PrintSlices(mark bool) {

	fmt.Print(pizza.SlicesAsString(mark))
}

func (pizza Pizza) PrintSlicesToFile(mark bool, path string) {

	text := pizza.SlicesAsString(mark)
	bytes := []byte(text)

	err := ioutil.WriteFile(path, bytes, 0644)
	simple.CheckErr(err)
}

func (pizza Pizza) Slices() []*Slice {

	tmp := make(map[string]*Slice)

	for _, xy := range pizza.Traversal() {
		cell := pizza.Cells[xy]
		sli := cell.Slice

		if sli != nil {
			tmp[sli.FormatCoordinates()] = sli
		}
	}

	slices := make([]*Slice, len(tmp))

	inx := 0
	for _, sli := range tmp {
		slices[inx] = sli
		inx++
	}

	return slices
}

func (pizza Pizza) SliceCount() int {

	return len(pizza.Slices())
}

func (pizza Pizza) Score() (covered int, score float32) {

	score = float32(pizza.ScoreTotal) / float32(pizza.Size())

	return pizza.ScoreTotal, score
}

func (pizza Pizza) PrintScore() {

	count, score := pizza.Score()

	fmt.Printf("Slices: %d\n", pizza.SliceCount())
	fmt.Printf("Covered cells: %d/%d\n", pizza.Size(), count)
	fmt.Printf("Percent: %.2f%%\n", score*100)
}

func (pizza Pizza) VectorPrint(row Vector, column Vector) {

	for _, iny := range row.Range() {
		for _, inx := range column.Range() {

			xy := Coordinate{Row: iny, Column: inx}
			cell := pizza.Cells[xy]
			fmt.Print(string(cell.Type))
		}

		fmt.Println()
	}
}

func (pizza Pizza) PrintPizza() {

	pizza.VectorPrint(pizza.Row, pizza.Column)
}

func (pizza Pizza) PrintVectors() {

	fmt.Printf("Row    := pizza.Vector{Start: %d, End: %d}\n", pizza.Row.Start, pizza.Row.End)
	fmt.Printf("Column := pizza.Vector{Start: %d, End: %d}\n", pizza.Column.Start, pizza.Column.End)
}

func (pizza Pizza) PrintSlicesVectors() {

	for _, sli := range pizza.Slices() {
		sli.PrintVector()
	}
}

func (pizza Pizza) PrintSlicesCoordinates() {

	for _, sli := range pizza.Slices() {
		sli.PrintCoordinates()
	}
}

func (pizza Pizza) submission() string {

	text := fmt.Sprint(pizza.SliceCount()) + "\n"

	for _, sli := range pizza.Slices() {

		text += sli.FormatCoordinates() + "\n"
	}

	return text
}

func (pizza Pizza) PrintSubmission() {
	fmt.Print(pizza.submission())
}

func (pizza Pizza) CreateSubmission(path string) {

	fmt.Print("Create submission ...\n")
	bytes := []byte(pizza.submission())
	err := ioutil.WriteFile(path, bytes, 0644)
	simple.CheckErr(err)
}

func (pizza *Pizza) AddSlice(slice *Slice) {

	for _, xy := range slice.Traversal() {
		cell := pizza.Cells[xy]

		if cell.Slice != nil {

			slice.PrintVector()
			cell.Slice.PrintVector()

			panic("Added overlaping slice...")
		}

		cell.Slice = slice
	}

	pizza.ScoreTotal += slice.Size()
}

func (pizza *Pizza) SafeAddSlice(slice *Slice) bool {

	for _, xy := range slice.Traversal() {
		cell := pizza.Cells[xy]

		if cell.Slice != nil {
			return false
		}
	}

	for _, xy := range slice.Traversal() {
		pizza.Cells[xy].Slice = slice
	}

	pizza.ScoreTotal += slice.Size()
	return true
}

func (pizza *Pizza) RemoveSlice(slice *Slice) {

	for _, xy := range slice.Traversal() {
		cell := pizza.Cells[xy]

		if cell.Slice != slice {
			panic("RemoveSlices wrong slice")
		}

		cell.Slice = nil
	}

	pizza.ScoreTotal -= slice.Size()
}

func (pizza *Pizza) RemoveAllSlice() {

	for _, xy := range pizza.Traversal() {
		pizza.Cells[xy].Slice = nil
	}

	pizza.ScoreTotal = 0
}

func (pizza Pizza) CheckErrors() {

	fmt.Println("Check for errors...")

	for _, sli1 := range pizza.Slices() {
		for _, sli2 := range pizza.Slices() {

			if sli1 == sli2 {
				continue
			}

			if sli1.Overlap(sli2) {
				fmt.Println("Overlap Error:")
				sli1.PrintVector()
				sli2.PrintVector()

				fmt.Println("Exit")
				simple.Exit()
			}
		}
	}
}

func (pizza Pizza) ContainsCoordinate(coordinate Coordinate) bool {

	xOk := pizza.Row.Start <= coordinate.Row && pizza.Row.End >= coordinate.Row
	yOk := pizza.Column.Start <= coordinate.Column && pizza.Column.End >= coordinate.Column

	return xOk && yOk
}

func (pizza Pizza) HasSliceAt(xy Coordinate) bool {

	return pizza.Cells[xy].Slice != nil
}
