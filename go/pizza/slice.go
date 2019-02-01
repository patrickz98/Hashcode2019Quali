package pizza

import "fmt"

type Slice struct {
	Pizza  *Pizza
	Row    Vector
	Column Vector
}

func (slice Slice) Size() int {

	return slice.Row.Length() * slice.Column.Length()
}

func (slice Slice) PrintVector() {
	fmt.Printf("row=%s column=%s\n", slice.Row.Stringify(), slice.Column.Stringify())
}

func (slice Slice) IngredientsOk() bool {

	// find.Pizza.PrintVector(rowV, cellV)
	tomato := 0
	mushroom := 0

	// fmt.Printf("row=%s cell=%s\n", rowV.Stringify(), cellV.Stringify())

	for _, iny := range slice.Row.Range() {
		for _, inx := range slice.Column.Range() {
			cell := slice.Pizza.Cells[ iny ][ inx ]

			// if cell.Slice != nil {
			// 	return false
			// }

			if cell.Type == 'T' {
				tomato++
			} else {
				mushroom++
			}
		}
	}

	ingredient := slice.Pizza.Ingredient

	return tomato >= ingredient && mushroom >= ingredient
}

func (slice Slice) Oversize() bool {

	return slice.Size() > slice.Pizza.MaxCells
}

func (slice Slice) HasMaxSize() bool {

	return slice.Size() == slice.Pizza.MaxCells
}

func (slice Slice) PrintInfo() {
	fmt.Printf("row: %s\n", slice.Row.Stringify())
	fmt.Printf("col: %s\n", slice.Column.Stringify())
	fmt.Printf("size: %d\n\n", slice.Size())
}

func (slice Slice) Overlap(slice2 *Slice) bool {

	row1 := slice.Row
	row2 := slice2.Row

	col1 := slice.Column
	col2 := slice2.Column

	return row1.Overlap(row2) && col1.Overlap(col2)
}

func (slice Slice) Equals(slice2 *Slice) bool {

	row1 := slice.Row
	row2 := slice2.Row

	col1 := slice.Column
	col2 := slice2.Column

	return row1.Equals(row2) && col1.Equals(col2)
}
