package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryExpand(xy pizza.Coordinate) (new Slices, overlap Slices) {

	bestGain := 0
	var newSlice *pizza.Slice
	var overlaps []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		candidateOverlaps := slicer.overlapSlices(sliceCandidate)

		destruction := 0

		for _, destructSlice := range candidateOverlaps {
			destruction += destructSlice.Size()
		}

		gain := sliceCandidate.Size() - destruction

		// sliceCandidate.VectorPrint()
		// fmt.Printf("gain=%d\n", gain)

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			overlaps = candidateOverlaps
		}
	}

	if newSlice == nil {
		return nil, nil
	}

	splitParts := slicer.splitSlice(newSlice)

	return splitParts, overlaps
}

func (slicer *Slicer) ExpandThroughDestruction() {

	fmt.Println("Expand through destruction")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {
		queue.Push(xy)
	}

	start, _ := slicer.Pizza.Score()

	for queue.HasItems() {

		fmt.Printf("Destruction queue --> %-7d\r", len(queue.data) - 1)
		// slicer.tryExpand(queue)

		xy := queue.Pop()

		slices, overlaps := slicer.tryExpand(*xy)

		for _, over := range overlaps {
			slicer.Pizza.RemoveSlice(over)
		}

		for _, slice := range slices {
			slicer.Pizza.AddSlice(slice)
		}
	}

	fmt.Println()

	now, _ := slicer.Pizza.Score()
	fmt.Printf("Destruction gain --> %d\n", now - start)
}
