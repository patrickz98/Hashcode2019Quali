package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryExpand(queue *CoordinateQueue) {

	xy := queue.Pop()

	// slicer.Pizza.PrintSlices()
	// fmt.Printf("xy=(%d, %d)\n", xy.Row, xy.Column)

	if slicer.Pizza.Cells[ *xy ].Slice != nil {
		return
	}

	bestGain := 0
	var newSlice *pizza.Slice
	var newSliceOverlaps []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ *xy ] {

		overlaps := slicer.overlapSlices(sliceCandidate)

		destruction := 0

		for _, destructSlice := range overlaps {
			destruction += destructSlice.Size()
		}

		gain := sliceCandidate.Size() - destruction

		// sliceCandidate.PrintVector()
		// fmt.Printf("gain=%d\n", gain)

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			newSliceOverlaps = overlaps
		}
	}

	if newSlice == nil {
		// queue.Push(*xy)
		return
	}

	newQueueElements := make(map[pizza.Coordinate] *pizza.Slice)

	for _, destructSlice := range newSliceOverlaps {
		slicer.Pizza.RemoveSlice(destructSlice)

		for _, xy := range destructSlice.Traversal() {
			newQueueElements[ xy ] = destructSlice
		}
	}

	for _, xy := range newSlice.Traversal() {
		delete(newQueueElements, xy);
	}

	splitParts := slicer.splitSlice(newSlice)

	for _, slice := range splitParts {
		slicer.Pizza.AddSlice(slice)
	}

	for xy := range newQueueElements {
		queue.Push(xy)
	}
}

func (slicer *Slicer) ExpandThroughDestruction() {

	fmt.Println("Expand through destruction")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		queue.Push(xy)
	}

	for queue.HasItems() {
		fmt.Printf("CoordinateQueue --> %-7d\r", len(queue.data))
		slicer.tryExpand(queue)
	}

	fmt.Println()
}
