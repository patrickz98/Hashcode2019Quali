package slicer

import (
	"../pizza"
	"../simple"
	"fmt"
	"os"
)

func overlap(pizzaa *pizza.PizzaPart, slice *pizza.Slice) bool {

	for _, sli := range pizzaa.Slices {

		if sli.Overlap(slice) {
			return true
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

	for _, sli := range part.Slices {

		if sli == orig {
			continue
		}

		if sli.Overlap(slice) {
			return true
		}
	}

	return false
}

func appendVector(array []*pizza.Vector, vector *pizza.Vector) []*pizza.Vector {

	for _, vec := range array {
		if vec.Equals(*vector) {
			return array
		}
	}

	return append(array, vector)
}

func vectorFind(vec pizza.Vector, sli pizza.Vector, maxExpand int) []*pizza.Vector {

	rVectors := make([]*pizza.Vector, 0)

	for rStart := 0; rStart <= maxExpand; rStart++ {
		for rEnd := 0; rEnd <= maxExpand; rEnd++ {

			start := simple.Max(vec.Start, sli.Start - rStart)
			end := simple.Min(vec.End, sli.End + rEnd)
			// start := sli.Start - rStart
			// end := sli.End + rEnd

			vec2 := &pizza.Vector{Start: start, End: end}

			if (vec2.Length() - vec.Length()) > maxExpand {
				continue
			}

			rVectors = appendVector(rVectors, vec2)
		}
	}

	return rVectors
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

		// part.PrintPart()
		// sli.Print()
		// fmt.Printf("rowExpand=%d\n", rowExpand)
		// fmt.Printf("colExpand=%d\n", colExpand)
		// fmt.Printf("sli.Row=%s\n", sli.Row.Stringify())
		// fmt.Printf("sli.Column=%s\n", sli.Column.Stringify())

		// TODO: Optimisation

		rowVectors := vectorFind(part.VectorR, sli.Row, rowExpand)
		colVetors := vectorFind(part.VectorC, sli.Column, colExpand)

		// fmt.Printf("part.VectorR=%s\n", part.VectorR.Stringify())
		// fmt.Printf("sli.Row=%s\n", sli.Row.Stringify())

		// for _, val := range rowVectors {
		// 	fmt.Printf("val=%s\n", val.Stringify())
		// }
		// fmt.Println()

		for _, row := range rowVectors {
			for _, col := range colVetors {

				expandedSlice := &pizza.Slice{
					Pizza: sli.Pizza,
					Row: *row,
					Column: *col,
				}

				// fmt.Printf("iny=%d inx=%d\n", iny, inx)
				// fmt.Printf("row=%s col=%s\n", row.Stringify(), col.Stringify())
				// expandedSlice.PrintInfo()
				// expandedSlice.PrintVector()

				if expandedSlice.Size() <= sli.Size() {
					continue
				}

				if expandedSlice.Oversize() {
					continue
				}

				if expandedSlice.Equals(sli) {
					continue
				}

				if ! expandedSlice.IngredientsOk() {
					continue
				}

				if checkOverlap(part, sli, expandedSlice) {
					continue
				}

				// fmt.Printf("row=%s\n", row.Stringify())
				// fmt.Printf("col=%s\n", col.Stringify())
				// fmt.Printf("size=%d\n", expandedSlice.Size())

				if expandedSlice.Size() > biggest.Size() {
					biggest = expandedSlice
				}

				// expandedSlice.PrintInfo()
			}
		}

		part.Slices[ sliceIndex ] = biggest

		// biggest.PrintInfo()
		// fmt.Println("**************************")
	}
}

func hasSliceAt(pizzaa *pizza.PizzaPart, iny int, inx int) bool {

	for _, sli := range pizzaa.Slices {
		row := sli.Row
		col := sli.Column

		if row.Start < iny && row.End > iny {
			return true
		}

		if col.Start < inx && col.End > inx {
			return true
		}
	}

	return false
}

func findAt(pizzaa *pizza.PizzaPart, iny int, inx int) {

	// fmt.Printf("find at iny=%d inx=%d\n", iny, inx)

	if hasSliceAt(pizzaa, iny, inx) {
		return
	}
	// fmt.Println(">>>>>>>> ok")

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
	mParts := make([]*pizza.PizzaPart, 0)

	for _, val := range parts {
		rpart := recursiveMatch(val)
		findSlices(rpart)
		expandSlices(rpart)

		mParts = append(mParts, rpart)
	}

	merged := merge(part.Pizza, mParts)

	// fmt.Println("part")
	// part.PrintSlices()
	// part.PrintScore()
	//
	// fmt.Println("merged")
	// merged.PrintSlices()
	// merged.PrintScore()

	findSlices(merged)
	expandSlices(merged)

	// _, _, score := merged.Score()
	total, count, score := merged.Score()
	fmt.Printf("total=%d count=%d score=%.2f\n", total, count, score * 100)

	if score > 1 {

		// mParts[3].PrintSlices()
		// mParts[3].PrintVectors()
		// for _, val := range parts {
		// 	val.PrintSlices()
		// }

		// fmt.Println()
		merged.PrintSlices()
		merged.PrintVectors()
		// merged.PrintPart()
		os.Exit(1)
	}

	return merged
}

func SearchSlices(pizz *pizza.Pizza) {

	start := pizza.InitPizzaPart(pizz)

	// start = &pizza.PizzaPart{
	// 	Pizza: pizz,
	// 	VectorR: pizza.Vector{Start: 13, End: 24},
	// 	VectorC: pizza.Vector{Start: 32, End: 47},
	// 	Slices: []*pizza.Slice{},
	// }

	test := recursiveMatch(start)
	test.PrintSlices()
	test.PrintScore()

	// findSlices(start)
	// expandSlices(start)
	//
	// start.PrintSlices()
	// start.PrintScore()

	// parts := start.Cut()[2].Cut()
	//
	// for inx, _ := range parts {
	// 	findSlices(parts[ inx ])
	// 	expandSlices(parts[ inx ])
	// }
	//
	// merged := merge(pizz, parts)
	// findSlices(merged)
	// expandSlices(merged)
	//
	// // merged.PrintSlices()
	//
	// parts2 := start.Cut()[3].Cut()
	//
	// for inx, _ := range parts2 {
	// 	findSlices(parts2[ inx ])
	// 	expandSlices(parts2[ inx ])
	// }
	//
	// merged2 := merge(pizz, parts2)
	// findSlices(merged2)
	// expandSlices(merged2)
	//
	// // merged2.PrintSlices()
	//
	// merged3 := merge(pizz, []*pizza.PizzaPart{merged, merged2})
	//
	// fmt.Println("######################")
	// merged3.PrintSlices()
	//
	// findSlices(merged3)
	// expandSlices(merged3)
	//
	// fmt.Println()
	// merged3.PrintSlices()


	// bab := start.Cut()
	//
	// part := bab[ 0 ]
	//
	// findSlices(part)
	// part.PrintSlices()
	//
	// fmt.Println("---------- expandSlices ----------")
	// expandSlices(part)
	// part.PrintSlices()


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
