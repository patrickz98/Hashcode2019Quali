package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryAllAt(queue *CoordinateQueue, xy pizza.Coordinate) {

	piz := slicer.Pizza

	if piz.HasSliceAt(xy) {
		return
	}

	neighbors := slicer.findBestNeighbor(xy)

	if neighbors != nil {
		slicer.AddSlices(neighbors.Slices)
		return
	}

	slices, overlaps := slicer.destructionAt(xy)

	if len(slices) > 0 {
		slicer.RemoveSlices(overlaps)
		slicer.AddSlices(slices)

		leftovers := slicer.leftovers(slices, overlaps)
		queue.PushAll(leftovers)

		return
	}

	slices, overlaps = slicer.shrinkAt(xy)

	if len(slices) > 0 {
		slicer.RemoveSlices(overlaps)
		slicer.AddSlices(slices)

		leftovers := slicer.leftovers(slices, overlaps)
		queue.PushAll(leftovers)

		return
	}

	// move, old := slicer.shakeAt(xy)
	//
	// if move != nil {
	// 	slicer.RemoveSlices(old)
	// 	slicer.AddSlices(move)
	//
	// 	leftovers := slicer.leftovers(move, old)
	// 	queue.PushAll(leftovers)
	//
	// 	return
	// }
}

func (slicer *Slicer) TryAll() {

	queue := InitCoordinateQueue()
	// queue.Push(pizza.Coordinate{Row: 0, Column: 0})
	queue.PushAll(slicer.Pizza.Traversal())

	// piz := slicer.Pizza
	// for _, xy := range piz.TraversalNotSlicedCells() {
	//
	// 	queue.Push(xy)
	// }

	for queue.HasItems() {

		xy := *queue.Pop()
		slicer.tryAllAt(queue, xy)

		// fmt.Printf("Expand queue --> %-7d\r", queue.Len())
		// fmt.Printf("%s\n", xy.Stringify())
	}

	fmt.Println()
}
