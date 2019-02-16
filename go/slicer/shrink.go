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

func (slicer *Slicer) tryExpandMove(queue *CoordinateQueue) {

	xy := queue.Pop()

	bestGain := 0
	var newSlice *pizza.Slice
	var sliceOverlaps []*pizza.Slice
	var sliceReplacements []*pizza.Slice
	// var leftovers []pizza.Coordinate

	for _, sliceCandidate := range slicer.SliceCache[ *xy ] {

		overlaps := slicer.overlapSlices(sliceCandidate)
		newSlices := make([]*pizza.Slice, 0)
		// newLeftovers := make([]pizza.Coordinate, 0)

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

			// for _, leftXY := range newSlice.Complement(shrinkSlice) {
			// 	newLeftovers = append(newLeftovers, leftXY)
			// }
		}

		if !replacementOk {
			break
		}

		gain := sliceCandidate.Size() - lost

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			sliceOverlaps = overlaps
			sliceReplacements = newSlices
			// leftovers = newLeftovers
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

	// for _, leftXY := range leftovers {
		// queue.Push(leftXY)
	// }
}

func (slicer *Slicer) ExpandThroughShrink() {

	fmt.Println("Expand through move")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		queue.Push(xy)
	}

	for queue.HasItems() {
		fmt.Printf("Move queue --> %-7d\r", len(queue.data) - 1)
		slicer.tryExpandMove(queue)
	}

	fmt.Printf("Move queue --> done\n")
}
