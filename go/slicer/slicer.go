package slicer

import (
	"../pizza"
	"../simple"
	"fmt"
)

func overlap(pizzaa *pizza.PizzaPart, slice *pizza.Slice) bool {

	for _, iny := range slice.Row.Range() {
		for _, inx := range slice.Column.Range() {
			cell := pizzaa.Pizza.Cells[ iny ][ inx ]

			if cell.Slice != nil {
				return true
			}
		}
	}

	return false
}

func findSlice(pizzaa *pizza.Pizza, row pizza.Vector, col pizza.Vector) []*pizza.Slice {

	max := pizzaa.MaxCells

	slices := make([]*pizza.Slice, 0)

	for _, iny := range row.Range() {
		for _, inx := range col.Range() {

			rowEnd := simple.Min(row.End, iny + max)
			searchR := pizza.Vector{Start:iny, End: rowEnd}

			colEnd := simple.Min(col.End, inx + max)
			searchC := pizza.Vector{Start:inx, End: colEnd}

			for _, r := range searchR.Range() {
				for _, c := range searchC.Range() {

					rowV := pizza.Vector{Start: iny, End: r}
					cellV := pizza.Vector{Start: inx, End: c}

					slic := &pizza.Slice{
						Pizza: pizzaa,
						Row: rowV,
						Column: cellV,
					}

					if slic.Oversize() {
						continue
					}

					if ! slic.IngredientsOk() {
						continue
					}

					// fmt.Printf("row: %s\n", rowV.Stringify())
					// fmt.Printf("cell: %s\n", cellV.Stringify())
					// fmt.Printf("size: %d\n", cellV.Size(rowV))

					slices = append(slices, slic)
				}
			}
		}
	}

	return slices
}

func checkOverlap(part *pizza.PizzaPart, orig *pizza.Slice, slice *pizza.Slice) bool {

	// fmt.Println("------------ check overlap ------------")
	// slice.PrintInfo()
	// fmt.Println("++++")

	for _, sli := range part.Slices {

		if sli == orig {
			continue
		}

		if sli.Overlap(slice) {

			// sli.PrintInfo()
			// fmt.Println("------------ overlap true ------------")
			return true
		}
	}

	// fmt.Println("------------ overlap false ------------")

	return false
}

func expandSlices(part *pizza.PizzaPart) {

	slices := part.Slices

	for sliceIndex, sli := range slices {
		if sli.HasMaxSize() {
			continue
		}

		// sliceSize := sli.Size()

		biggest := sli

		rowExpand := (part.Pizza.MaxCells / sli.Column.Length()) - sli.Row.Length()
		colExpand := (part.Pizza.MaxCells / sli.Row.Length())    - sli.Column.Length()

		sli.PrintInfo()
		fmt.Printf("rowExpand=%d\n", rowExpand)
		fmt.Printf("colExpand=%d\n", colExpand)

		// TODO: Optimisation

		for iny := -rowExpand; iny <= rowExpand; iny++ {

			var row pizza.Vector

			if iny < 0 {
				rStart := simple.Max(part.VectorR.Start, sli.Row.Start + iny)
				row = pizza.Vector{Start: rStart, End: sli.Row.End}
			} else {
				rEnd := simple.Min(part.VectorR.End, sli.Row.End + iny)
				row = pizza.Vector{Start: sli.Row.Start, End: rEnd}
			}

			if iny == 0 {
				row = sli.Row
			}

			for inx := -colExpand; inx <= colExpand; inx++ {

				var col pizza.Vector

				if inx < 0 {
					cStart := simple.Max(part.VectorC.Start, sli.Column.Start + inx)
					col = pizza.Vector{Start: cStart, End: sli.Column.End}
				} else {
					cEnd := simple.Min(part.VectorC.End, sli.Column.End+inx)
					col = pizza.Vector{Start: sli.Column.Start, End: cEnd}
				}

				if inx == 0 {
					col = sli.Column
				}

				expandedSlice := &pizza.Slice{
					Pizza: sli.Pizza,
					Row: row,
					Column: col,
				}

				// fmt.Printf("iny=%d inx=%d\n", iny, inx)
				// fmt.Printf("row=%s col=%s\n", row.Stringify(), col.Stringify())
				// expandedSlice.PrintInfo()
				// expandedSlice.PrintVector()

				if expandedSlice.Equals(sli) {
					continue
				}

				if ! expandedSlice.IngredientsOk() {
					continue
				}

				if checkOverlap(part, sli, expandedSlice) {
					continue
				}

				if biggest.Size() < expandedSlice.Size() {
					biggest = expandedSlice
				}

				// expandedSlice.PrintInfo()
			}
		}

		part.Slices[ sliceIndex ] = biggest

		biggest.PrintInfo()
		fmt.Println("**************************")
	}
}

func findAt(pizzaa *pizza.PizzaPart, iny int, inx int) {

	cell := pizzaa.Pizza.Cells[ iny ][ inx ]

	if cell.Slice != nil {
		return
	}

	max := pizzaa.Pizza.MaxCells

	rowEnd := simple.Min(pizzaa.VectorR.End, iny + max)
	row := pizza.Vector{Start:iny, End: rowEnd}

	colEnd := simple.Min(pizzaa.VectorC.End, inx + max)
	col := pizza.Vector{Start:inx, End: colEnd}

	slices := findSlice(pizzaa.Pizza, row, col)

	var smallest *pizza.Slice

	for _, slic := range slices {

		if slic == nil {
			continue
		}

		if overlap(pizzaa, slic) {
			continue
		}

		// slic.PrintInfo()

		if (smallest == nil) || (smallest.Size() > slic.Size()) {
			smallest = slic
		}
	}

	if smallest != nil {
		pizzaa.AddSlice(*smallest)
	}
}

func findSlices(part *pizza.PizzaPart) {

	// findAt(part, 0, 0)

	for _, iny := range part.VectorR.Range() {
		for _, inx := range part.VectorC.Range() {
			findAt(part, iny, inx)
		}
	}
}

func SearchSlices(pizz *pizza.Pizza) {

	start := pizza.InitPizzaPart(pizz)

	// test := recursiveMatch(start)
	// test.PrintSlices()
	// test.PrintScore()

	// findSlices(start)
	// start.PrintSlices()
	// start.PrintScore()

	bab := start.Cut()

	part := bab[ 0 ]

	findSlices(part)
	part.PrintSlices()

	fmt.Println("---------- expandSlices ----------")
	expandSlices(part)

	part.PrintSlices()

	// start.Slices = part.Slices
	// start.PrintSlices()

	// for inx := range parts {
	// 	findSlices(parts[ inx ])
	//
	// 	fmt.Println("-----------------")
	// 	parts[ inx ].PrintSlices()
	// 	fmt.Println("-----------------")
	// }
	//
	// test := merge(start.Pizza, parts)
	// test.PrintSlices()

	// parts2 := parts[ 3 ].Cut()
	//
	// findSlices(parts2[ 3 ])
	//
	// parts2[ 3 ].PrintPart()
	// parts2[ 3 ].PrintSlices()
}

func merge(pizz *pizza.Pizza, parts []*pizza.PizzaPart) *pizza.PizzaPart {

	slices := make([]*pizza.Slice, 0)

	var rVector *pizza.Vector
	var cVector *pizza.Vector

	for _, part := range parts {

		if rVector == nil {
			rVector = &part.VectorR
		} else {
			rVector = rVector.Join(part.VectorR)
		}

		if cVector == nil {
			cVector = &part.VectorC
		} else {
			cVector = cVector.Join(part.VectorC)
		}

		for _, sli := range part.Slices {
			slices = append(slices, sli)
		}
	}

	return &pizza.PizzaPart{
		Pizza: pizz,
		Slices: slices,
		VectorR: *rVector,
		VectorC: *cVector,
	}
}

func recursiveMatch(part *pizza.PizzaPart) *pizza.PizzaPart {

	if ! part.CutPossible() {

		return part
	}

	parts := part.Cut()

	for inx, val := range parts {
		parts[ inx ] = recursiveMatch(val)
		findSlices(parts[ inx ])

		// fmt.Println("-----------------")
		// parts[ inx ].PrintSlices()
		// fmt.Println("-----------------")
	}

	merged := merge(part.Pizza, parts)

	// Fill open gaps
	findSlices(merged)

	return merged
}
