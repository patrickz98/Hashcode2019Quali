package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryDestuctShinkAt(xy pizza.Coordinate) (slices Slices, overlaps Slices) {

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
		return nil , nil
	}

	if distGain > shrinkGain {
		return distSlices, distOverlaps
	} else {
		return shrinkSlices, shrinkOverlaps
	}
}

func (slicer *Slicer) TryDestuctShink() {

	queue := InitCoordinateQueue()
	// queue.Push(pizza.Coordinate{Row: 0, Column: 0})
	// queue.PushAll(slicer.Pizza.Traversal())
	queue.PushAll(slicer.Pizza.TraversalNotSlicedCells())

	score, _ := slicer.Pizza.Score()

	for queue.HasItems() {

		xy := *queue.Pop()

		if slicer.Pizza.HasSliceAt(xy) {
			continue
		}

		slices, overlaps := slicer.tryDestuctShinkAt(xy)

		if slices == nil {
			continue
		}

		slicer.RemoveSlices(overlaps)
		slicer.AddSlices(slices)

		leftovers := slicer.leftovers(slices, overlaps)
		queue.PushAll(leftovers)

		gain, _ := slicer.Pizza.Score()
		fmt.Printf("Expand gain=%d queue=%-7d\r", gain - score, queue.Len())
	}

	fmt.Println()
}
