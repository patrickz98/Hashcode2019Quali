package slicer

import (
	"fmt"
)

func (slicer *Slicer) TryAll() {

	fmt.Println("Expand...")

	piz := slicer.Pizza

	queue := InitCoordinateQueue()

	for _, xy := range piz.TraversalNotSlicedCells() {

		queue.Push(xy)
	}

	for queue.HasItems() {
		// xy := *queue.PopFist()
		xy := *queue.Pop()

		// fmt.Println(xy.Stringify())
		// simple.Exit()

		fmt.Printf("Expand queue --> %-7d\r", queue.Len())

		if piz.HasSliceAt(xy) {
			continue
		}

		neighbors := slicer.findBestNeighbor(xy)

		if neighbors != nil {
			slicer.AddSlices(neighbors.Slices)
			continue
		}

		// smallest := slicer.findSmallestAt(xy)
		//
		// if smallest != nil {
		// 	piz.AddSlice(smallest)
		// 	continue
		// }

		slices, overlaps := slicer.destructionAt(xy)

		if len(slices) > 0 {
			slicer.RemoveSlices(overlaps)
			slicer.AddSlices(slices)

			leftovers := slicer.leftovers(slices, overlaps)

			for _, leftXY := range leftovers {
				queue.Push(leftXY)
			}

			continue
		}

		slices, overlaps = slicer.shrinkAt(xy)

		if len(slices) > 0 {
			slicer.RemoveSlices(overlaps)
			slicer.AddSlices(slices)

			leftovers := slicer.leftovers(slices, overlaps)

			for _, leftXY := range leftovers {
				queue.Push(leftXY)
			}

			continue
		}

		// move, old := slicer.tryMove(xy)
		//
		// if move == nil || old == nil {
		// 	continue
		// }
		//
		// // fmt.Println("Move")
		//
		// slicer.RemoveSlice(old)
		// slicer.AddSlice(move)
		//
		// leftovers := slicer.leftover(move, old)
		//
		// for _, leftXY := range leftovers {
		// 	queue.Push(leftXY)
		// }
	}

	fmt.Println()
}
