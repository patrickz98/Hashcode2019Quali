package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) forceExpand(xy pizza.Coordinate) {

	if slicer.Pizza.Cells[ xy ].Slice != nil {
		return
	}

	var forceSlice *pizza.Slice
	var overlaps []*pizza.Slice

	slices := slicer.SliceCache[ xy ]

	for _, slice := range slices {

		if slice == nil {
			continue
		}

		if (forceSlice == nil) || (forceSlice.Size() > slice.Size()) {
			forceSlice = slice
		}
	}

	if forceSlice == nil {
		return
	}

	for _, sli := range overlaps {
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
