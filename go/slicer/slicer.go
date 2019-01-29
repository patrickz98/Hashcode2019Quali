package slicer

import (
	"../pizza"
)

func valid(pizzaa *pizza.Pizza, rowV pizza.Vector, cellV pizza.Vector) bool {

	// find.Pizza.PrintVector(rowV, cellV)
	ingredient := pizzaa.Ingredient

	tomato := 0
	mushroom := 0

	// fmt.Printf("row=%s cellV=%s\n", rowV.Stringify(), cellV.Stringify())

	for iny := range rowV.Range() {
		for inx := range cellV.Range() {
			run := pizzaa.Cells[ iny ][ inx ].Type

			if run == 'T' {
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

	var biggest *pizza.Slice

	for r := iny; r <  iny + max; r++ {
		for c := inx; c < inx + max; c++ {
			// fmt.Printf("(%d, %d)\n", r, c)

			rowV := pizza.Vector{Start: iny, End: r}
			cellV := pizza.Vector{Start: inx, End: c}

			if pizzaa.Columns.End < cellV.End || pizzaa.Rows.End < rowV.End {
				continue
			}

			// for iny := rowV.Start; iny < rowV.End+1; iny++ {
			// 	columns := find.Pizza.Cells[ iny ][ cellV.Start : cellV.End+1 ]
			//
			// 	for _, cell := range columns {
			// 		if cell.Slice != nil {
			// 			continue
			// 		}
			// 	}
			// }

			if rowV.Size(cellV) > max {
				continue
			}

			// fmt.Printf("row: %s\n", rowV.Stringify())
			// fmt.Printf("cell: %s\n", cellV.Stringify())
			// fmt.Printf("size: %d\n", cellV.Size(rowV))

			if ! valid(pizzaa, rowV, cellV) {
				continue
			}

			slic := pizza.Slice{Row: rowV, Column: cellV}

			if (biggest == nil) || (biggest.Size() < slic.Size()) {
				biggest = &slic
			}

			// find.Pizza.Slices = make([]pizza.Slice, 1)
			// find.Pizza.Slices[ 0 ] =

			// find.Pizza.PrintSlices()
		}
	}

	if biggest != nil {
		// fmt.Println("Biggest")
		// biggest.PrintVector()
		// find.Pizza.PrintSlice(*biggest)
		// fmt.Println("-------")

		pizzaa.AddSlice(*biggest)

		// size := len(find.Pizza.Slices)
		// find.Pizza.Slices[ size ] = *biggest
	}
}

func FindSlice(part *pizza.Pizza) {

	// find.find(0, 0)

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
