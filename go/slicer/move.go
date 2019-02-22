package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryMove(xy pizza.Coordinate) {

	if slicer.Pizza.Cells[ xy ].Slice != nil {
		return
	}

	slices := slicer.SliceCache[ xy ]

	for _, moveCandidate := range slices {

		overlaps := slicer.overlapSlices(moveCandidate)

		if len(overlaps) > 1 {
			continue
		}

		for _, overlap := range overlaps {

			if moveCandidate == overlap {
				continue
			}

			if moveCandidate.Size() != overlap.Size() {
				continue
			}

			slicer.Pizza.RemoveSlice(overlap)
			slicer.Pizza.AddSlice(moveCandidate)

			return
		}
	}
}

func (slicer *Slicer) MoveSlices() {

	fmt.Println("Move existing slices")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {
		queue.Push(xy)
	}

	for queue.HasItems() {
		fmt.Printf("Move queue --> %-7d\r", len(queue.data) - 1)
		slicer.tryMove(*queue.Pop())
	}

	fmt.Printf("Move queue --> done\n")
}
