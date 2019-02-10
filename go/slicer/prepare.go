package slicer

import (
	"../pizza"
	"../simple"
	"fmt"
)

type Slicer struct {
	Pizza *pizza.PizzaPart
	Slices map[pizza.Coordinate] []*pizza.Slice
}

func (slicer *Slicer) Init() {
	slicer.buildSlicesCache()
}

func (slicer *Slicer) buildSlicesCache() {

	max := slicer.Pizza.Pizza.MaxCells

	slices := make(map[pizza.Coordinate][]*pizza.Slice)

	total := slicer.Pizza.Size()

	slicesCount := 0

	for count, coordinate := range slicer.Pizza.Traversal() {

		rowEnd := simple.Min(slicer.Pizza.VectorR.End, coordinate.Row+max)
		searchR := pizza.Vector{Start: coordinate.Row, End: rowEnd}

		colEnd := simple.Min(slicer.Pizza.VectorC.End, coordinate.Column+max)
		searchC := pizza.Vector{Start: coordinate.Column, End: colEnd}

		// Test all possible Slice dimensions.
		for _, endR := range searchR.Range() {
			for _, endC := range searchC.Range() {

				rowV := pizza.Vector{Start: coordinate.Row, End: endR}
				cellV := pizza.Vector{Start: coordinate.Column, End: endC}

				slic := &pizza.Slice{
					Pizza:  slicer.Pizza.Pizza,
					Row:    rowV,
					Column: cellV,
				}

				if ! slic.Ok() {
					continue
				}

				// Add Slice to each x and y position.
				for _, xy := range slic.Traversal() {
					slices[ xy ] = append(slices[ xy ], slic)
				}

				slicesCount++
			}
		}

		// fmt.Printf("Generating possible slices: %3.0f%%\r", (float32(count) / float32(total) * 100.0))
		fmt.Printf("Generating possible slices %d/%d\r", total, count + 1)
	}

	fmt.Println()
	// fmt.Printf("Generating possible slices: Done\n")
	fmt.Printf("Generated %d slices\n", slicesCount)

	slicer.Slices = slices
}

func (slicer *Slicer) overlap(slice *pizza.Slice) bool {

	// TODO: Optimise
	for _, sli := range slicer.Pizza.Slices {

		if sli.Overlap(slice) {
			return true
		}
	}

	return false
}

func (slicer *Slicer) FindSmallestParts() {

	for _, xy := range slicer.Pizza.Traversal() {

		slices := slicer.Slices[ xy ]

		var smallest *pizza.Slice

		for _, slice := range slices {

			if slice == nil {
				continue
			}

			if slicer.overlap(slice) {
				continue
			}

			// slic.PrintInfo()

			if (smallest == nil) || (smallest.Size() > slice.Size()) {
				smallest = slice
			}
		}

		if smallest != nil {
			slicer.Pizza.AddSlice(*smallest)
		}
	}
}