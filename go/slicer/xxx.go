package slicer

import (
	"../pizza"
	"fmt"
)

// var dist = 0
// var shrink = 0

func (slicer *Slicer) tryAllAt(queue *CoordinateQueue, xy pizza.Coordinate) {

	neighbors := slicer.findBestNeighbor(xy)

	if neighbors != nil {
		slicer.AddSlices(neighbors.Slices)
		return
	}

	distSlices, distOverlaps := slicer.destructionAt(xy)
	shrinkSlices, shrinkOverlaps := slicer.shrinkAt(xy)

	distGain := slicer.CalculateGain(distSlices, distOverlaps)
	shrinkGain := slicer.CalculateGain(shrinkSlices, shrinkOverlaps)

	if distGain == 0 && shrinkGain == 0 {
		return
	}

	// fmt.Printf("distGain=%d shrinkGain=%d\n", distGain, shrinkGain)

	if distGain > shrinkGain {
		// dist++
		slicer.RemoveSlices(distOverlaps)
		slicer.AddSlices(distSlices)

		leftovers := slicer.leftovers(distSlices, distOverlaps)
		queue.PushAll(leftovers)

	} else {
		// shrink++
		slicer.RemoveSlices(shrinkOverlaps)
		slicer.AddSlices(shrinkSlices)

		leftovers := slicer.leftovers(shrinkSlices, shrinkOverlaps)
		queue.PushAll(leftovers)
	}

	// fmt.Printf("dist=%d shrink=%d\n", dist, shrink)

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

	// queue.PushStart(xy)
}

func (slicer *Slicer) TryAll() {

	queue := InitCoordinateQueue()
	// queue.Push(pizza.Coordinate{Row: 0, Column: 0})
	// queue.PushAll(slicer.Pizza.Traversal())
	queue.PushAll(slicer.Pizza.TraversalNotSlicedCells())

	score, _ := slicer.Pizza.Score()

	for queue.HasItems() {

		xy := *queue.Pop()
		slicer.tryAllAt(queue, xy)

		gain, _ := slicer.Pizza.Score()
		fmt.Printf("Expand gain=%d queue=%-7d\r", gain - score, queue.Len())

		// fmt.Printf("%s\n", xy.Stringify())
	}

	fmt.Println()
}
