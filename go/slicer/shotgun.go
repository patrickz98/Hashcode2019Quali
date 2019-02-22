package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) forceExpand(xy pizza.Coordinate) {

	if slicer.Pizza.Cells[ xy ].Slice != nil {
		return
	}

	good := 0
	var forceSlice *pizza.Slice
	var forceOverlaps []*pizza.Slice

	slices := slicer.SliceCache[ xy ]

	for _, slice := range slices {

		overlaps := slicer.overlapSlices(slice)
		size := slice.Size()
		destruction := slicer.CalculateSize(overlaps)

		gain := size + destruction

		if forceSlice == nil || good < gain {
			good = gain
			forceSlice = slice
			forceOverlaps = overlaps
		}
	}

	if forceSlice == nil {
		return
	}

	for _, sli := range forceOverlaps {
		slicer.Pizza.RemoveSlice(sli)
	}

	slicer.Pizza.AddSlice(forceSlice)
}

func (slicer *Slicer) ExpandShot() {

	fmt.Println("Expand force...")

	for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {
		slicer.forceExpand(xy)
	}
}
