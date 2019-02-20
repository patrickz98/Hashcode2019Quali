package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) calcNeighborFactor(slice *pizza.Slice) float32 {

	rowStart := slice.Row.Start - 1
	rowEnd := slice.Row.End + 1

	colStart := slice.Column.Start - 1
	colEnd := slice.Column.End + 1

	factor := 0

	for iny := rowStart; iny <= rowEnd; iny++ {
		for inx := colStart; inx <= colEnd; inx++ {

			xy := pizza.Coordinate{Row: iny, Column: inx}

			if slice.ContainsCoordinate(xy) {
				continue
			}

			if !slicer.Pizza.ContainsCoordinate(xy) {
				factor++
				continue
			}

			if slicer.Pizza.HasSliceAt(xy) {
				factor++
			}
		}
	}

	return float32(factor) / float32(slice.Size())
}

func (slicer *Slicer) findBestFitEdge(xy pizza.Coordinate) {

	slices := slicer.SliceCache[ xy ]

	bestNeighborFactor := float32(0)
	var slice *pizza.Slice

	for _, sli := range slices {

		if slicer.overlap(sli) {
			continue
		}

		neighborFactor := slicer.calcNeighborFactor(sli)

		if bestNeighborFactor < neighborFactor {
			bestNeighborFactor = neighborFactor
			slice = sli
		}
	}

	if slice == nil {

		// slicer.tryExpand(xy)
		// slicer.tryExpandShrink(xy)
		return
	}

	// fmt.Printf("bestNeighborFactor --> %d\n", bestNeighborFactor)
	// slice.Print()
	// slice.PrintVector()

	// simple.Exit()

	slicer.Pizza.AddSlice(slice)
}

func (slicer *Slicer) ExpandThroughNeighbors() {

	fmt.Println("Expand edge...")

	for _, xy := range slicer.Pizza.Traversal() {

		if !slicer.Pizza.HasSliceAt(xy) {
			slicer.findBestFitEdge(xy)
		}
	}
}
