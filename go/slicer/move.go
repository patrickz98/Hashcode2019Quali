package slicer

import (
	"../pizza"
	"fmt"
	// "os"
)

func (slicer *Slicer) shrinkSlice(trigger *pizza.Slice, shrink *pizza.Slice) *pizza.Slice {

	// TODO: Add support for cutting slices.
	// fmt.Println("Trigger:")
	// trigger.Print()

	// fmt.Println("Shrink:")
	// shrink.Print()

	parts := slicer.slicesInSlice(shrink)

	var replacement *pizza.Slice

	for _, part := range parts {

		if part.Overlap(trigger) {
			continue
		}

		if (replacement == nil) || (replacement.Size() < part.Size()) {
			replacement = part
		}
	}

	// fmt.Println("Successor:")
	// replacement.Print()

	return replacement
}

func (slicer *Slicer) tryExpandMove(xy pizza.Coordinate) {

	bestGain := 0
	var newSlice *pizza.Slice
	var sliceOverlaps []*pizza.Slice
	var sliceReplacements []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		overlaps := slicer.overlapSlices(sliceCandidate)
		newSlices := make([]*pizza.Slice, len(overlaps))

		lost := 0

		replacementsFund := true

		for inx, shrinkSlice := range overlaps {

			newSlice := slicer.shrinkSlice(sliceCandidate, shrinkSlice)

			if newSlice == nil {
				replacementsFund = false
				break
			}

			// fmt.Printf("shrinkSlice=%d newSlice=%d\n", shrinkSlice.Size(), newSlice.Size())

			lost += shrinkSlice.Size() - newSlice.Size()
			newSlices[ inx ] = newSlice
		}

		if !replacementsFund {
			break
		}

		// if lost == 0 {
		// 	continue
		// }

		gain := sliceCandidate.Size() - lost
		// fmt.Printf("gain=%d lost=%d\n", sliceCandidate.Size(), lost)

		// sliceCandidate.PrintVector()
		// fmt.Printf("gain=%d\n", gain)

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			sliceOverlaps = overlaps
			sliceReplacements = newSlices
		}
	}

	if newSlice == nil {
		return
	}

	for _, slice := range sliceOverlaps {
		slicer.Pizza.RemoveSlice(slice)
	}

	for _, slice := range sliceReplacements {
		slicer.Pizza.AddSlice(slice)
	}

	slicer.Pizza.AddSlice(newSlice)
}

func (slicer *Slicer) ExpandThroughMove() {

	fmt.Println("Expand through move")

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		fmt.Printf("(%4d, %4d)\r", xy.Row, xy.Column)

		slicer.tryExpandMove(xy)
	}
}
