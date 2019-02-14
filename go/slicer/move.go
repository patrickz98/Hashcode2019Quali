package slicer

import (
	"../pizza"
	"fmt"
	// "os"
)

type shrinkStatus int

const (
	failed  shrinkStatus = -1
	success shrinkStatus = 0
	eaten   shrinkStatus = 1
)

func (slicer *Slicer) shrinkSlice(trigger *pizza.Slice, shrink *pizza.Slice) (shrinkStatus, *pizza.Slice) {

	// TODO: Add support for cutting slices.

	if trigger.Contains(shrink) {
		return eaten, nil
	}

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

	if replacement == nil {
		return failed, replacement
	} else {
		return success, replacement
	}
}

func (slicer *Slicer) tryExpandMove(xy pizza.Coordinate) {

	bestGain := 0
	var newSlice *pizza.Slice
	var sliceOverlaps []*pizza.Slice
	var sliceReplacements []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		overlaps := slicer.overlapSlices(sliceCandidate)
		newSlices := make([]*pizza.Slice, 0)

		lost := 0
		replacementOk := true

		for _, shrinkSlice := range overlaps {

			status, newSlice := slicer.shrinkSlice(sliceCandidate, shrinkSlice)

			if status == failed {
				replacementOk = false
				break
			}

			if status == eaten {
				continue
			}

			lost += shrinkSlice.Size() - newSlice.Size()
			newSlices = append(newSlices, newSlice)
		}

		if !replacementOk {
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

	for inx, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		fmt.Printf("Try to move %d/%d\r", slicer.Pizza.Size(), inx + 1)

		slicer.tryExpandMove(xy)
	}

	fmt.Println()
}
