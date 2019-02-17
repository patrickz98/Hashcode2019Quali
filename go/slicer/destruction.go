package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryExpand(xy pizza.Coordinate) {

	// xy := queue.Pop()

	// slicer.Pizza.PrintSlices()
	// fmt.Printf("xy=(%d, %d)\n", xy.Row, xy.Column)

	// if slicer.Pizza.Cells[ xy ].Slice != nil {
	// 	return
	// }

	bestGain := 0
	var newSlice *pizza.Slice
	var overlaps []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		posoverlaps := slicer.overlapSlices(sliceCandidate)

		destruction := 0

		for _, destructSlice := range posoverlaps {
			destruction += destructSlice.Size()
		}

		gain := sliceCandidate.Size() - destruction

		// sliceCandidate.VectorPrint()
		// fmt.Printf("gain=%d\n", gain)

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			overlaps = posoverlaps
		}
	}

	if newSlice == nil {
		return
	}

	for _, over := range overlaps {
		slicer.Pizza.RemoveSlice(over)
	}

	splitParts := slicer.splitSlice(newSlice)

	for _, slice := range splitParts {
		slicer.Pizza.AddSlice(slice)
	}
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
		slicer.tryExpand(*xy)
	}

	fmt.Println()

	now, _ := slicer.Pizza.Score()
	fmt.Printf("Destruction gain --> %d\n", now - start)

	// fmt.Println()
	// fmt.Printf("Destruction queue --> done\n")
}
