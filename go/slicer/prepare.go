package slicer

import (
	"../pizza"
	"../simple"
	"fmt"
)

type Slicer struct {
	Pizza *pizza.Pizza
	Slices map[pizza.Coordinate] []*pizza.Slice
}

func (slicer *Slicer) Init() {
	slicer.buildSlicesCache()
}

func (slicer *Slicer) buildSlicesCache() {

	max := slicer.Pizza.MaxCells

	slices := make(map[pizza.Coordinate][]*pizza.Slice)

	total := slicer.Pizza.Size()

	slicesCount := 0

	for count, coordinate := range slicer.Pizza.Traversal() {

		rowEnd := simple.Min(slicer.Pizza.Row.End, coordinate.Row+max)
		searchR := pizza.Vector{Start: coordinate.Row, End: rowEnd}

		colEnd := simple.Min(slicer.Pizza.Column.End, coordinate.Column+max)
		searchC := pizza.Vector{Start: coordinate.Column, End: colEnd}

		// Test all possible Slice dimensions.
		for _, endR := range searchR.Range() {
			for _, endC := range searchC.Range() {

				rowV := pizza.Vector{Start: coordinate.Row, End: endR}
				cellV := pizza.Vector{Start: coordinate.Column, End: endC}

				slic := &pizza.Slice{
					Pizza:  slicer.Pizza,
					Row:    rowV,
					Column: cellV,
				}

				if ! slic.Ok() {
					continue
				}

				// Add Slice to each x and y position.
				for _, xy := range slic.Traversal() {
					slices[ xy ] = append(slices[ xy ], slic)
				}

				slicesCount++
			}
		}

		// fmt.Printf("Generating possible slices: %3.0f%%\r", (float32(count) / float32(total) * 100.0))
		fmt.Printf("Generating possible slices %d/%d\r", total, count + 1)
	}

	fmt.Println()
	// fmt.Printf("Generating possible slices: Done\n")
	fmt.Printf("Generated %d slices\n", slicesCount)

	slicer.Slices = slices
}

func (slicer *Slicer) overlap(slice *pizza.Slice) bool {

	// TODO: Optimise
	for _, xy := range slice.Traversal() {

		cell := slicer.Pizza.Cells[ xy ]
		if cell.Slice != nil {
			return true
		}
	}

	return false
}

func (slicer *Slicer) overlapSlices(slice *pizza.Slice) []*pizza.Slice {

	overlap := make([]*pizza.Slice, 0)

	for _, xy := range slice.Traversal() {

		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			overlap = append(overlap, cell.Slice)
		}
	}

	return overlap
}

func (slicer *Slicer) tryExpand(queue *CoordinateQueue) {
// func (slicer *Slicer) tryExpand(queue *chan pizza.Coordinate) {
// func (slicer *Slicer) tryExpand(queue *[]pizza.Coordinate) {

	// list := *queue
	// xy := <-*queue
	xy := queue.Pop()

	bestGain := 0
	var newSlice *pizza.Slice
	var newSliceOverlaps []*pizza.Slice

	for _, sliceCandidate := range slicer.Slices[ *xy ] {
		overlaps := slicer.overlapSlices(sliceCandidate)

		destruction := 0

		for _, destructSlice := range overlaps {
			destruction += destructSlice.Size()
		}

		gain := sliceCandidate.Size() - destruction

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			newSliceOverlaps = overlaps
		}
	}

	if newSlice == nil {
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

	slicer.Pizza.AddSlice(newSlice)

	for xy := range newQueueElements {
		queue.Push(xy)
		// *queue <- xy
		// tmp := append(*queue, xy)
		// queue = &tmp
	}
}

func (slicer *Slicer) ExpandThroughDestruction() {

	fmt.Println("ExpandThroughDestruction")

	// queue := make([]pizza.Coordinate, 0)
	// queue := make(chan pizza.Coordinate)
	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		// queue = append(queue, xy)
		// queue <- xy
		queue.Push(xy)
	}

	for queue.HasItems() {
		// fmt.Printf("len(queue)=%d\n", len(queue))
		// fmt.Printf("len(queue)=%d\n", len(queue.data))
		slicer.tryExpand(queue)
	}
}

func (slicer *Slicer) FindSmallestParts() {

	size := slicer.Pizza.Size()

	for count, xy := range slicer.Pizza.Traversal() {

		slices := slicer.Slices[ xy ]

		var smallest *pizza.Slice

		for _, slice := range slices {

			if slice == nil {
				continue
			}

			if slicer.overlap(slice) {
				continue
			}

			// slic.PrintInfo()

			if (smallest == nil) || (smallest.Size() > slice.Size()) {
				smallest = slice
			}
		}

		if smallest != nil {
			slicer.Pizza.AddSlice(smallest)
		}

		fmt.Printf("Find smalest slices %d/%d\r", size, count + 1)
	}

	fmt.Println()
}