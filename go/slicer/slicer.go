package slicer

import (
	"../pizza"
	"../simple"
)

func valid(pizzaa *pizza.PizzaPart, rowV pizza.Vector, cellV pizza.Vector) bool {

	// find.Pizza.PrintVector(rowV, cellV)
	tomato := 0
	mushroom := 0

	// fmt.Printf("row=%s cell=%s\n", rowV.Stringify(), cellV.Stringify())

	for _, iny := range rowV.Range() {
		for _, inx := range cellV.Range() {
			cell := pizzaa.Pizza.Cells[ iny ][ inx ]

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

	ingredient := pizzaa.Pizza.Ingredient

	return tomato >= ingredient && mushroom >= ingredient
}

func find(pizzaa *pizza.PizzaPart, iny int, inx int) {

	cell := pizzaa.Pizza.Cells[ iny ][ inx ]

	if cell.Slice != nil {
		return
	}

	max := pizzaa.Pizza.MaxCells

	rowEnd := simple.Min(pizzaa.VectorR.End, iny + max)
	row := pizza.Vector{Start:iny, End: rowEnd}

	colEnd := simple.Min(pizzaa.VectorC.End, inx + max)
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

func findSlices(part *pizza.PizzaPart) {

	// find(part, 0, 0)

	for _, iny := range part.VectorR.Range() {
		for _, inx := range part.VectorC.Range() {
			find(part, iny, inx)
		}
	}
}

func SearchSlices(pizz *pizza.Pizza) {

	start := pizza.InitPizzaPart(pizz)

	// test := recursiveMatch(start)
	// test.PrintSlices()
	// test.PrintScore()

	findSlices(start)
	start.PrintSlices()
	start.PrintScore()


	// bab := start.Cut()
	// parts := bab[ 1 ].Cut()
	// bab[ 1 ].PrintSlices()
	//
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

func expandSlices(part *pizza.PizzaPart) {


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
