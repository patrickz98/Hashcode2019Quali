package pizza

import (
	"fmt"
)

type Slice struct {
	Pizza  *Pizza
	Row    Vector
	Column Vector
}

func (slice Slice) Size() int {

	return slice.Row.Length() * slice.Column.Length()
}

func (slice Slice) IngredientsOk() bool {

	tomato := 0
	mushroom := 0

	for _, xy := range slice.Traversal() {
		cell := slice.Pizza.Cells[ xy ]

		if cell.Type == 'T' {
			tomato++
		} else {
			mushroom++
		}
	}

	ingredients := slice.Pizza.Ingredients

	return tomato >= ingredients && mushroom >= ingredients
}

func (slice Slice) Oversize() bool {

	return slice.Size() > slice.Pizza.MaxCells
}

func (slice Slice) Valid() bool {

	return !slice.Oversize() && slice.IngredientsOk()
}

func (slice Slice) Overlap(slice2 *Slice) bool {

	row1 := slice.Row
	row2 := slice2.Row

	col1 := slice.Column
	col2 := slice2.Column

	overlay := row1.Overlap(row2) && col1.Overlap(col2)

	return overlay
}

func (slice Slice) Equals(slice2 *Slice) bool {

	row1 := slice.Row
	row2 := slice2.Row

	col1 := slice.Column
	col2 := slice2.Column

	return row1.Equals(row2) && col1.Equals(col2)
}

func (slice Slice) Traversal() []Coordinate {

	coordinates := make([]Coordinate, slice.Size())

	for iny, row := range slice.Row.Range() {
		for inx, col := range slice.Column.Range() {
			index := (iny * slice.Column.Length()) + inx
			coordinates[ index ] = Coordinate{Row: row, Column: col}
		}
	}

	return coordinates
}

func (slice Slice) Contains(slice2 *Slice) bool {

	return slice.Row.ContainsVector(slice2.Row) && slice.Column.ContainsVector(slice2.Column)
}

func (slice Slice) ContainsCoordinate(coordinate Coordinate) bool {

	xOk := slice.Row.Start <= coordinate.Row && slice.Row.End >= coordinate.Row
	yOk := slice.Column.Start <= coordinate.Column && slice.Column.End >= coordinate.Column

	return xOk && yOk
}

func (slice Slice) Complement(slice2 *Slice) []Coordinate {

	complement := make([]Coordinate, 0)

	for _, xy := range slice2.Traversal() {

		if !slice.ContainsCoordinate(xy) {
			complement = append(complement, xy)
		}
	}

	return complement
}

func (slice Slice) Print() {

	slice.Pizza.VectorPrint(slice.Row, slice.Column)
}

func (slice Slice) PrintInfo() {
	fmt.Printf("row: %s\n", slice.Row.Stringify())
	fmt.Printf("col: %s\n", slice.Column.Stringify())
	fmt.Printf("size: %d\n\n", slice.Size())
}

func (slice Slice) FormatCoordinates() string {

	format := fmt.Sprintf("%d %d %d %d", slice.Row.Start, slice.Column.Start,
		slice.Row.End, slice.Column.End)

	return format
}

func (slice Slice) FormatVectors() string {

	return fmt.Sprintf("row%s column%s", slice.Row.Stringify(), slice.Column.Stringify())
}

func (slice Slice) PrintVector() {
	fmt.Println(slice.FormatVectors())
}

func (slice Slice) PrintCoordinates() {
	fmt.Println(slice.FormatCoordinates())
}
