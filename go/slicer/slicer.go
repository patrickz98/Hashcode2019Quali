package slicer

import (
	"../pizza"
	"../simple"
)

// func overlap(pizzaa *pizza.PizzaPart, slice *pizza.Slice) bool {
//
// 	for _, sli := range pizzaa.Slices {
//
// 		if sli.Overlap(slice) {
// 			return true
// 		}
// 	}
//
// 	return false
// }
//
// func overlapSlices(pizzaa *pizza.PizzaPart, slice *pizza.Slice) []*pizza.Slice {
//
// 	overlap := make([]*pizza.Slice, 0)
//
// 	for _, sli := range pizzaa.Slices {
//
// 		if sli.Overlap(slice) {
// 			overlap = append(overlap, sli)
// 		}
// 	}
//
// 	return overlap
// }
//
// func findSlice(pizzaa *pizza.Pizza, row pizza.Vector, col pizza.Vector) []*pizza.Slice {
//
// 	max := pizzaa.MaxCells
//
// 	slices := make([]*pizza.Slice, 0)
//
// 	for _, iny := range row.Range() {
// 		for _, inx := range col.Range() {
//
// 			rowEnd := simple.Min(row.End, iny + max)
// 			searchR := pizza.Vector{Start:iny, End: rowEnd}
//
// 			colEnd := simple.Min(col.End, inx + max)
// 			searchC := pizza.Vector{Start:inx, End: colEnd}
//
// 			for _, r := range searchR.Range() {
// 				for _, c := range searchC.Range() {
//
// 					rowV := pizza.Vector{Start: iny, End: r}
// 					cellV := pizza.Vector{Start: inx, End: c}
//
// 					slic := &pizza.Slice{
// 						Pizza: pizzaa,
// 						Row: rowV,
// 						Column: cellV,
// 					}
//
// 					if slic.Ok() {
// 						slices = append(slices, slic)
// 					}
// 				}
// 			}
// 		}
// 	}
//
// 	return slices
// }

// func checkOverlap(part *pizza.PizzaPart, orig *pizza.Slice, slice *pizza.Slice) bool {
//
// 	for _, sli := range part.Slices {
//
// 		if sli == orig {
// 			continue
// 		}
//
// 		if sli.Overlap(slice) {
// 			return true
// 		}
// 	}
//
// 	return false
// }

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

// func expandSlices(part *pizza.PizzaPart) {
//
// 	slices := part.Slices
//
// 	for sliceIndex, sli := range slices {
// 		if sli.HasMaxSize() {
// 			continue
// 		}
//
// 		// sliceSize := sli.Size()
//
// 		biggest := sli
//
// 		rowExpand := (part.Pizza.MaxCells / sli.Column.Length()) - sli.Row.Length()
// 		colExpand := (part.Pizza.MaxCells / sli.Row.Length())    - sli.Column.Length()
//
// 		// part.PrintPart()
// 		// sli.Print()
// 		// fmt.Printf("rowExpand=%d\n", rowExpand)
// 		// fmt.Printf("colExpand=%d\n", colExpand)
// 		// fmt.Printf("sli.Row=%s\n", sli.Row.Stringify())
// 		// fmt.Printf("sli.Column=%s\n", sli.Column.Stringify())
//
// 		// TODO: Optimisation
//
// 		rowVectors := vectorFind(part.VectorR, sli.Row, rowExpand)
// 		colVetors := vectorFind(part.VectorC, sli.Column, colExpand)
//
// 		// fmt.Printf("part.VectorR=%s\n", part.VectorR.Stringify())
// 		// fmt.Printf("sli.Row=%s\n", sli.Row.Stringify())
//
// 		// for _, val := range rowVectors {
// 		// 	fmt.Printf("val=%s\n", val.Stringify())
// 		// }
// 		// fmt.Println()
//
// 		for _, row := range rowVectors {
// 			for _, col := range colVetors {
//
// 				expandedSlice := &pizza.Slice{
// 					Pizza: sli.Pizza,
// 					Row: *row,
// 					Column: *col,
// 				}
//
// 				// fmt.Printf("iny=%d inx=%d\n", iny, inx)
// 				// fmt.Printf("row=%s col=%s\n", row.Stringify(), col.Stringify())
// 				// expandedSlice.PrintInfo()
// 				// expandedSlice.PrintVector()
//
// 				if expandedSlice.Size() <= sli.Size() {
// 					continue
// 				}
//
// 				if expandedSlice.Oversize() {
// 					continue
// 				}
//
// 				if expandedSlice.Equals(sli) {
// 					continue
// 				}
//
// 				if ! expandedSlice.IngredientsOk() {
// 					continue
// 				}
//
// 				if checkOverlap(part, sli, expandedSlice) {
// 					continue
// 				}
//
// 				// fmt.Printf("row=%s\n", row.Stringify())
// 				// fmt.Printf("col=%s\n", col.Stringify())
// 				// fmt.Printf("size=%d\n", expandedSlice.Size())
//
// 				if expandedSlice.Size() > biggest.Size() {
// 					biggest = expandedSlice
// 				}
//
// 				// expandedSlice.PrintInfo()
// 			}
// 		}
//
// 		part.Slices[ sliceIndex ] = biggest
//
// 		// biggest.PrintInfo()
// 		// fmt.Println("**************************")
// 	}
// }

// func hasSliceAt(pizzaa *pizza.PizzaPart, iny int, inx int) bool {
//
// 	for _, sli := range pizzaa.Slices {
// 		row := sli.Row
// 		col := sli.Column
//
// 		rowMatch := row.Start <= iny && row.End >= iny
// 		colMatch := col.Start <= inx && col.End >= inx
//
// 		if rowMatch && colMatch {
// 			return true
// 		}
// 	}
//
// 	return false
// }

// func findAt(part *pizza.PizzaPart, iny int, inx int) {
//
// 	// fmt.Printf("find at iny=%d inx=%d\n", iny, inx)
//
// 	if hasSliceAt(part, iny, inx) {
// 		return
// 	}
// 	// fmt.Println(">>>>>>>> ok")
//
// 	max := part.Pizza.MaxCells
//
// 	rowEnd := simple.Min(part.VectorR.End, iny + max)
// 	row := pizza.Vector{Start:iny, End: rowEnd}
//
// 	colEnd := simple.Min(part.VectorC.End, inx + max)
// 	col := pizza.Vector{Start:inx, End: colEnd}
//
// 	slices := findSlice(part.Pizza, row, col)
//
// 	var smallest *pizza.Slice
//
// 	for _, slic := range slices {
//
// 		if slic == nil {
// 			continue
// 		}
//
// 		if overlap(part, slic) {
// 			continue
// 		}
//
// 		// slic.PrintInfo()
//
// 		if (smallest == nil) || (smallest.Size() > slic.Size()) {
// 			smallest = slic
// 		}
// 	}
//
// 	if smallest != nil {
// 		part.AddSlice(*smallest)
// 	}
// }

// func findSlices(part *pizza.PizzaPart) {
//
// 	// findAt(part, 0, 0)
//
// 	for _, iny := range part.VectorR.Range() {
// 		for _, inx := range part.VectorC.Range() {
// 			findAt(part, iny, inx)
// 		}
// 	}
// }

// func merge(pizz *pizza.Pizza, parts []*pizza.PizzaPart) *pizza.PizzaPart {
//
// 	slices := make([]*pizza.Slice, 0)
//
// 	var rVector *pizza.Vector
// 	var cVector *pizza.Vector
//
// 	for _, part := range parts {
//
// 		if rVector == nil {
// 			rVector = &part.VectorR
// 		} else {
// 			rVector = rVector.Join(part.VectorR)
// 		}
//
// 		if cVector == nil {
// 			cVector = &part.VectorC
// 		} else {
// 			cVector = cVector.Join(part.VectorC)
// 		}
//
// 		for _, sli := range part.Slices {
// 			slices = append(slices, sli)
// 		}
// 	}
//
// 	return &pizza.PizzaPart{
// 		Pizza: pizz,
// 		Slices: slices,
// 		VectorR: *rVector,
// 		VectorC: *cVector,
// 	}
// }
//
// func deleteSlice(part *pizza.PizzaPart, slice *pizza.Slice) {
//
// 	index := -1
//
// 	for inx, sli := range part.Slices {
// 		if sli == slice {
// 			index = inx
// 		}
// 	}
//
// 	if index < 0 {
// 		return
// 	}
//
// 	part.Slices = append(part.Slices[ index + 1: ], part.Slices[ : index]...)
// }

// func findNewByBreak(part *pizza.PizzaPart) {
//
// 	// fmt.Println("-------------------------------")
// 	// part.PrintSlices()
//
// 	total, count, _ := part.Score()
// 	if total == count {
// 		return
// 	}
//
// 	max := part.Pizza.MaxCells
//
// 	newSlices := make(map[*pizza.Slice] bool)
//
// 	for _, iny := range part.VectorR.Range() {
// 		for _, inx := range part.VectorC.Range() {
//
// 			if hasSliceAt(part, iny, inx) {
// 				// fmt.Printf("hasSliceAt(%d, %d)\n", iny, inx)
// 				continue
// 			}
//
// 			// TODO: DO in all directions
// 			rowEnd := simple.Min(part.VectorR.End, iny + max)
// 			row := pizza.Vector{Start:iny, End: rowEnd}
//
// 			colEnd := simple.Min(part.VectorC.End, inx + max)
// 			col := pizza.Vector{Start:inx, End: colEnd}
//
// 			slices := findSlice(part.Pizza, row, col)
//
// 			bigSmallReplacements := make(map[*pizza.Slice] []*pizza.Slice)
//
// 			for _, sli := range slices {
//
// 				// fmt.Printf("iny=%d inx=%d\n", iny, inx)
// 				// sli.Print()
// 				// sli.PrintVector()
//
// 				overlap := overlapSlices(part, sli)
//
// 				lostSize := 0
//
// 				for _, over := range overlap {
//
// 					_, changed := newSlices[ over ]
//
// 					if !changed {
// 						lostSize += over.Size()
// 					}
// 				}
//
// 				// fmt.Printf("overlap=%d\n", lostSize)
// 				// fmt.Printf("sli.Size=%d\n", sli.Size())
//
// 				if lostSize <= sli.Size() {
//
// 					bigSmallReplacements[ sli ] = overlap
// 				}
// 			}
//
// 			var smallest * pizza.Slice
//
// 			for repl, _ := range bigSmallReplacements {
// 				// fmt.Println("Repl")
// 				// repl.Print()
//
// 				if smallest == nil {
// 					smallest = repl
// 					continue
// 				}
//
// 				if (repl.Size() < smallest.Size()) && !repl.Equals(smallest) {
// 					smallest = repl
// 				}
// 			}
//
// 			if smallest == nil {
// 				continue
// 			}
//
// 			// fmt.Println("Smallest")
// 			// smallest.Print()
//
// 			for _, key := range bigSmallReplacements[ smallest ] {
//
// 				// fmt.Println("Delete")
// 				// key.Print()
//
// 				deleteSlice(part, key)
// 			}
//
// 			part.Slices = append(part.Slices, smallest)
// 			newSlices[ smallest ] = true
// 		}
// 	}
//
// 	// part.PrintSlices()
// 	//
// 	// fmt.Println("+++++++++++++++++++++++++++++++")
// }

// func recursiveMatch(part *pizza.PizzaPart) *pizza.PizzaPart {
//
// 	if ! part.CutPossible() {
//
// 		return part
// 	}
//
// 	parts := part.Cut()
// 	mParts := make([]*pizza.PizzaPart, 0)
//
// 	for _, val := range parts {
// 		rpart := recursiveMatch(val)
// 		findSlices(rpart)
// 		expandSlices(rpart)
// 		// findNewByBreak(rpart)
// 		// findSlices(rpart)
// 		// expandSlices(rpart)
//
// 		mParts = append(mParts, rpart)
// 	}
//
// 	merged := merge(part.Pizza, mParts)
//
// 	// fmt.Println("part")
// 	// part.PrintSlices()
// 	// part.PrintScore()
// 	//
// 	// fmt.Println("merged")
// 	// merged.PrintSlices()
// 	// merged.PrintScore()
//
// 	findSlices(merged)
// 	expandSlices(merged)
// 	// findNewByBreak(merged)
// 	// findSlices(merged)
// 	// expandSlices(merged)
//
// 	// _, _, score := merged.Score()
// 	total, count, score := merged.Score()
// 	percent := float32(total) / (float32(part.Pizza.Columns) * float32(part.Pizza.Rows))
//
// 	// fmt.Printf("total=%5d count=%5d score=%6.2f\n", total, count, score * 100)
// 	fmt.Printf("(%6.2f%%) total=%5d count=%5d score=%6.2f\n", percent, total, count, score * 100)
//
// 	if score > 1 {
//
// 		// mParts[3].PrintSlices()
// 		// mParts[3].PrintVectors()
// 		// for _, val := range parts {
// 		// 	val.PrintSlices()
// 		// }
//
// 		// fmt.Println()
// 		merged.PrintSlices()
// 		merged.PrintVectors()
// 		merged.PrintSlicesPlain()
// 		// merged.PrintPart()
// 		os.Exit(1)
// 	}
//
// 	return merged
// }

func SearchSlices(pizza *pizza.Pizza) {

	slicer := Slicer{Pizza: pizza}

	slicer.Init()
	// slicer.FindBiggestParts()
	// slicer.FindSmallestParts()
	slicer.ExpandThroughDestruction()

	pizza.PrintSlices()
	pizza.PrintScore()

	// pizza.PrintSlicesToFile("xxx.txt")

	// start = &pizza.PizzaPart{
	// 	Pizza: pizz,
	// 	VectorR: pizza.Vector{Start: 13, End: 24},
	// 	VectorC: pizza.Vector{Start: 32, End: 47},
	// 	Slices: []*pizza.Slice{},
	// }

	// test := recursiveMatch(start)
	// test.PrintSlices()
	// test.PrintScore()

	// findSlices(start)
	// findNewByBreak(start)
	// expandSlices(start)

	// start.PrintSlices()

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
}
