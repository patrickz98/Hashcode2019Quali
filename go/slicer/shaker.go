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

func (slicer *Slicer) ShakeSlicesWithQueue(queue *CoordinateQueue) {

	changed := 0

	//fmt.Printf("--> Shake changed=%d queue=%-7d \r", changed, queue.Len())

	for queue.HasItems() {

		xy := *queue.Pop()

		//fmt.Printf("--> Shake changed=%d queue=%-7d \r", changed, queue.Len())

		slices, old := slicer.shakeAt(xy)

		if slices == nil {
			continue
		}

		changed++

		slicer.RemoveSlices(old)
		slicer.AddSlices(slices)
	}

	//fmt.Println()
}

func (slicer *Slicer) ShakeSlices() {

	fmt.Println("Shake existing slices")

	queue := InitCoordinateQueue()
	queue.PushAll(slicer.Pizza.Traversal())
	//queue.PushAll(slicer.Pizza.TraversalNotSlicedCells())

	slicer.ShakeSlicesWithQueue(queue)
}
