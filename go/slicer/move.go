package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) tryMove(xy pizza.Coordinate) (move *pizza.Slice, old *pizza.Slice) {

	slices := slicer.SliceCache[ xy ]

	for _, moveCandidate := range slices {

		overlaps := slicer.overlapSlices(moveCandidate)

		if len(overlaps) != 1 {
			continue
		}

		overlap := overlaps[ 0 ]

		if moveCandidate == overlap {
			continue
		}

		if moveCandidate.Size() != overlap.Size() {
			continue
		}

		return moveCandidate, overlap
	}

	return nil, nil
}

func (slicer *Slicer) MoveSlices() {

	fmt.Println("Move existing slices")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {
		queue.Push(xy)
	}

	for queue.HasItems() {
		fmt.Printf("Move queue --> %-7d\r", queue.Len())

		xy := *queue.Pop()

		if slicer.Pizza.Cells[ xy ].Slice != nil {
			return
		}

		moved, old := slicer.tryMove(xy)

		if moved == nil || old == nil {
			continue
		}

		fmt.Println("---------- old ----------")
		old.PrintVector()
		old.Print()
		slicer.Pizza.RemoveSlice(old)

		fmt.Println("---------- moved ----------")
		moved.PrintVector()
		moved.Print()
		slicer.Pizza.AddSlice(moved)
	}

	fmt.Printf("Move queue --> done\n")
}
