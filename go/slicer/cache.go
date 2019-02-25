package slicer

import (
	"../pizza"
	"../simple"
	"fmt"
	"sort"
)

func (slicer *Slicer) buildSlicesCache() {

	max := slicer.Pizza.MaxCells

	slices := make(map[ pizza.Coordinate ] Slices)
	slicer.TopLeftSliceCache = make(map[ pizza.Coordinate ] Slices)

	total := slicer.Pizza.Size()

	slicesCount := 0

	for count, coordinate := range slicer.Pizza.Traversal() {

		rowEnd := simple.Min(slicer.Pizza.Row.End, coordinate.Row+max)
		searchR := pizza.Vector{Start: coordinate.Row, End: rowEnd}

		colEnd := simple.Min(slicer.Pizza.Column.End, coordinate.Column+max)
		searchC := pizza.Vector{Start: coordinate.Column, End: colEnd}

		// Test all possible Slice dimensions.
		for _, endR := range searchR.Range() {
			for _, endC := range searchC.Range() {

				rowV := pizza.Vector{Start: coordinate.Row, End: endR}
				cellV := pizza.Vector{Start: coordinate.Column, End: endC}

				slic := &pizza.Slice{
					Pizza:  slicer.Pizza,
					Row:    rowV,
					Column: cellV,
				}

				if !slic.Valid() {
					continue
				}

				slicer.TopLeftSliceCache[coordinate] = append(slicer.TopLeftSliceCache[coordinate], slic)

				// Add Slice to each x and y position.
				for _, xy := range slic.Traversal() {
					slices[xy] = append(slices[xy], slic)
				}

				slicesCount++
			}
		}

		// fmt.Printf("Generating possible slices: %3.0f%%\r", (float32(count) / float32(total) * 100.0))
		fmt.Printf("Generating possible slices %d/%d\r", total, count+1)
	}

	fmt.Println()
	// fmt.Printf("Generating possible slices: Done\n")
	fmt.Printf("Generated %d slices\n", slicesCount)

	slicer.SliceCache = slices

	for key := range slices {
		sort.Slice(slices[key], func(i int, j int) bool {
			return slices[key][i].Size() < slices[key][j].Size()
		})
	}
}
