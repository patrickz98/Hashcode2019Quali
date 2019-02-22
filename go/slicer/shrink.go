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

func (slicer *Slicer) shrinkCutSlice(trigger *pizza.Slice, shrink *pizza.Slice) (shrinkStatus, int, []*pizza.Slice) {

	// TODO: Not working.

	if trigger.Contains(shrink) {
		return eaten, 0, nil
	}

	parts := slicer.slicesInSlice(shrink)
	replacements := make([]*pizza.Slice, 0)

	for _, part := range parts {

		if part.Overlap(trigger) {
			continue
		}

		replacements = append(replacements, part)
	}

	if len(replacements) <= 0 {
		return failed, 0, nil
	}

	bestSum := 0
	bestReplacements := make([]*pizza.Slice, 0)

	for _, set := range slicer.powerSet(replacements) {

		overlap := false
		sum := 0

		for _, sli1 := range set {

			for _, sli2 := range set {

				if !overlap && sli1 != sli2 && sli1.Overlap(sli2) {
					overlap = true
				}
			}

			sum += sli1.Size()
		}

		if overlap {
			continue
		}

		if bestSum < sum && len(bestReplacements) < len(set) {
			bestSum = sum
			bestReplacements = set
		}
	}

	if len(bestReplacements) == 0 {
		return failed, 0, nil
	}

	return success, bestSum, bestReplacements
}

func (slicer *Slicer) shrinkSlice(trigger *pizza.Slice, shrink *pizza.Slice) (shrinkStatus, *pizza.Slice) {

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

func (slicer *Slicer) tryExpandShrink(xy pizza.Coordinate) {

	bestGain := 0
	var newSlice *pizza.Slice
	var sliceOverlaps []*pizza.Slice
	var sliceReplacements []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		overlaps := slicer.overlapSlices(sliceCandidate)
		newSlices := make([]*pizza.Slice, 0)

		lostSize := 0
		replaceSize := 0
		replacementOk := true

		for _, shrinkSlice := range overlaps {

			status, newSlice := slicer.shrinkSlice(sliceCandidate, shrinkSlice)

			if status == failed {
				replacementOk = false
				break
			}

			lostSize += sliceCandidate.Size()

			if status == eaten {
				continue
			}

			replaceSize += newSlice.Size()
			newSlices = append(newSlices, newSlice)
		}

		if !replacementOk {
			continue
		}

		gain := replaceSize + sliceCandidate.Size() - lostSize

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

	splitParts := slicer.splitSlice(newSlice)

	for _, slice := range splitParts {
		slicer.Pizza.AddSlice(slice)
	}
}

func (slicer *Slicer) ExpandThroughShrink() {

	fmt.Println("Expand through shrink")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {
		queue.Push(xy)
	}

	start, _ := slicer.Pizza.Score()

	for queue.HasItems() {

		fmt.Printf("Shrink queue --> %-7d\r", len(queue.data) - 1)

		xy := queue.Pop()
		slicer.tryExpandShrink(*xy)
	}

	fmt.Println()

	now, _ := slicer.Pizza.Score()
	fmt.Printf("Shrink gain --> %d\n", now - start)

	// fmt.Println()
	// fmt.Printf("Shrink queue --> done\n")
}
