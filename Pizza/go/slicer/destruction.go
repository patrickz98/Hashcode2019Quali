package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) destructionAt(xy pizza.Coordinate) (new Slices, overlap Slices) {

	bestGain := 0
	var newSlice *pizza.Slice
	var overlaps []*pizza.Slice

	for _, sliceCandidate := range slicer.SliceCache[ xy ] {

		candidateOverlaps := slicer.overlapSlices(sliceCandidate)

		destruction := 0

		for _, destructSlice := range candidateOverlaps {
			destruction += destructSlice.Size()
		}

		gain := sliceCandidate.Size() - destruction

		if bestGain < gain {
			bestGain = gain
			newSlice = sliceCandidate
			overlaps = candidateOverlaps
		}
	}

	if newSlice == nil {
		return nil, nil
	}

	splitParts := slicer.splitSlice(newSlice)

	return splitParts, overlaps
}

func (slicer *Slicer) ExpandThroughDestruction() {

	fmt.Println("Expand through destruction")

	queue := InitCoordinateQueue()
	queue.PushAll(slicer.Pizza.Traversal())

	start, _ := slicer.Pizza.Score()

	for queue.HasItems() {

		fmt.Printf("Destruction queue --> %-7d\r", len(queue.data) - 1)
		// slicer.destructionAt(queue)

		xy := queue.Pop()

		slices, overlaps := slicer.destructionAt(*xy)

		slicer.RemoveSlices(overlaps)
		slicer.AddSlices(slices)

		leftovers := slicer.leftovers(slices, overlaps)
		queue.PushAll(leftovers)
	}

	fmt.Println()

	now, _ := slicer.Pizza.Score()
	fmt.Printf("Destruction gain --> %d\n", now - start)
}

func (slicer *Slicer) ExpandThroughDestructionBrute() {

	fmt.Println("Expand through destruction")

	start, _ := slicer.Pizza.Score()

	count := 0
	covered := 0

	for {
		fmt.Printf("Round %d covered=%d\n", count, covered)

		bestGain := 0
		var bestSlices Slices
		var bestOverlap Slices

		for _, xy := range slicer.Pizza.Traversal() {

			slis, overs := slicer.destructionAt(xy)

			if slis == nil {
				continue
			}

			gain := slicer.CalculateSize(slis) - slicer.CalculateSize(overs)

			if bestGain < gain {
				fmt.Printf("%s\n", xy.Stringify())

				bestGain = gain
				bestSlices = slis
				bestOverlap = overs
				break
			}
		}

		count++

		if bestSlices != nil {

			slicer.RemoveSlices(bestOverlap)
			slicer.AddSlices(bestSlices)
			covered += bestGain

			continue
		}

		break
	}

	fmt.Println()

	now, _ := slicer.Pizza.Score()
	fmt.Printf("Destruction gain --> %d\n", now - start)
}
