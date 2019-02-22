package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) destructionAt(xy pizza.Coordinate) (new Slices, overlap Slices) {

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
		// slicer.destructionAt(queue)

		xy := queue.Pop()

		slices, overlaps := slicer.destructionAt(*xy)

		slicer.RemoveSlices(overlaps)
		slicer.AddSlices(slices)
	}

	fmt.Println()

	now, _ := slicer.Pizza.Score()
	fmt.Printf("Destruction gain --> %d\n", now - start)
}
