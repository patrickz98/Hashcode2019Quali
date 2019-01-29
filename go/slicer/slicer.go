package slicer

import (
	"../pizza"
	"../simple"
)

func valid(pizzaa *pizza.Pizza, rowV pizza.Vector, cellV pizza.Vector) bool {

	// find.Pizza.PrintVector(rowV, cellV)
	ingredient := pizzaa.Ingredient

	tomato := 0
	mushroom := 0

	// fmt.Printf("row=%s cell=%s\n", rowV.Stringify(), cellV.Stringify())

	for _, iny := range rowV.Range() {
		for _, inx := range cellV.Range() {
			cell := pizzaa.Cells[ iny ][ inx ]

			if cell.Slice != nil {
				return false
			}

			if cell.Type == 'T' {
				tomato++
			} else {
				mushroom++
			}
		}
	}

	return tomato >= ingredient && mushroom >= ingredient
}

func find(pizzaa *pizza.Pizza, iny int, inx int) {

	cell := pizzaa.Cells[ iny ][ inx ]

	if cell.Slice != nil {
		return
	}

	max := pizzaa.MaxCells

	rowEnd := simple.Min(pizzaa.Rows.End, iny + max)
	row := pizza.Vector{Start:iny, End: rowEnd}

	colEnd := simple.Min(pizzaa.Columns.End, inx + max)
	col := pizza.Vector{Start:inx, End: colEnd}

	// fmt.Printf("### iny=%d inx=%d\n", iny, inx)
	// fmt.Printf("### row=%s col=%s\n", row.Stringify(), col.Stringify())

	var biggest *pizza.Slice

	for _, r := range row.Range() {
		for _, c := range col.Range() {

			rowV := pizza.Vector{Start: iny, End: r}
			cellV := pizza.Vector{Start: inx, End: c}

			slic := pizza.Slice{Row: rowV, Column: cellV}

			if slic.Size() > max {
				continue
			}

			// fmt.Printf("row: %s\n", rowV.Stringify())
			// fmt.Printf("cell: %s\n", cellV.Stringify())
			// fmt.Printf("size: %d\n", cellV.Size(rowV))

			if ! valid(pizzaa, rowV, cellV) {
				continue
			}


			if (biggest == nil) || (biggest.Size() < slic.Size()) {
				biggest = &slic
			}
		}
	}

	if biggest != nil {
		// fmt.Println("Biggest")
		// biggest.PrintVector()
		// find.Pizza.PrintSlice(*biggest)
		// fmt.Println("-------")

		pizzaa.AddSlice(*biggest)
	}
}

func FindSlice(part *pizza.Pizza) {

	// find(part, 0, 0)

	for _, iny := range part.Rows.Range() {
		for _, inx := range part.Columns.Range() {
			find(part, iny, inx)
		}
	}
}

func cut(part *pizza.Pizza) {

	if ! part.CutPossible() {
		return
	}

	parts := part.Cut()

	for _, val := range parts {

		if val != nil {
			cut(val)


			// merge

			break
		}
	}
}
