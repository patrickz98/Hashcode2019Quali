package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) shakeAt(xy pizza.Coordinate) (new Slices, overlap Slices) {

	// bestGain := 0
	var newSlice *pizza.Slice
	var overlaps []*pizza.Slice

	sliceAtXY := slicer.Pizza.Cells[ xy ].Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		if sliceAtXY == sliceCandidate {
			continue
		}
		
		candidateOverlaps := slicer.overlapSlices(sliceCandidate)

		destruction := slicer.CalculateSize(candidateOverlaps)
		gain := sliceCandidate.Size() - destruction

		if gain == 0 {
			// bestGain = gain
			newSlice = sliceCandidate
			overlaps = candidateOverlaps
			break
		}
	}

	if newSlice == nil {
		return nil, nil
	}

	// newSlice.Print()
	// fmt.Printf("fund=%d\n", len(slices))

	splitParts := slicer.splitSlice(newSlice)

	return splitParts, overlaps
}

func (slicer *Slicer) ShakeSlices() {

	fmt.Println("Change existing slices")

	queue := InitCoordinateQueue()

	// for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {
	for _, xy := range slicer.Pizza.Traversal() {
		queue.Push(xy)
	}

	changed := 0

	for queue.HasItems() {
		fmt.Printf("--> Change changed=%d queue=%-7d \r", changed, queue.Len())

		xy := *queue.Pop()

		// if slicer.Pizza.HasSliceAt(xy) {
		// 	return
		// }

		slices, old := slicer.shakeAt(xy)

		if slices == nil {
			continue
		}

		changed++

		// fmt.Printf("fund=%d\n", len(slices))

		slicer.RemoveSlices(old)
		slicer.AddSlices(slices)
	}

	fmt.Println()
	// fmt.Printf("Change queue --> done\n")
}
