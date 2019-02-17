package slicer

import "fmt"
// import "../simple"
import "../pizza"

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

			// fmt.Printf("trigger: (%d, %d)\n", xy.Row, xy.Column)
			//
			// fmt.Println("moveCandidate:")
			// moveCandidate.Print()
			// moveCandidate.VectorPrint()
			//
			// fmt.Println("overlap:")
			// overlap.Print()
			// overlap.VectorPrint()
			//
			// for _, xxy := range moveCandidate.Complement(overlap) {
			// 	fmt.Printf("(%d, %d)\n", xxy.Row, xxy.Column)
			// }
			//
			// simple.Exit()
		}
	}
}

func (slicer *Slicer) ExpandThroughMove() {

	fmt.Println("Expand through move")

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
