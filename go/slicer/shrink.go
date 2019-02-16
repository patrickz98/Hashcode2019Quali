package slicer

import (
	"../pizza"
	"fmt"
	// "os"
)

type shrinkStatus int

const (
	failed  shrinkStatus = -1
	success shrinkStatus = 0
	eaten   shrinkStatus = 1
)

func (slicer *Slicer) shrinkSlice(trigger *pizza.Slice, shrink *pizza.Slice) (shrinkStatus, int, []*pizza.Slice) {

	if trigger.Contains(shrink) {
		return eaten, 0, nil
	}

	parts := slicer.slicesInSlice(shrink)

	replacements := make([]*pizza.Slice, 0)

	for _, part := range parts {

		if part.Overlap(trigger) {
			continue
		}

		replacements = append(replacements, part)
	}

	if len(replacements) <= 0 {
		return failed, 0, nil
	}

	bestSum := 0
	var bestReplacements []*pizza.Slice

	for _, set := range slicer.powerSet(replacements) {

		overlap := false
		sum := 0

		for _, sli1 := range set {

			for _, sli2 := range set {

				if sli1 != sli2 && sli1.Overlap(sli2) {
					overlap = true
				}
			}

			sum += sli1.Size()
		}

		if overlap {
			continue
		}

		if bestSum > sum {
			continue
		}

		bestReplacements = set
	}

	// if len(bestReplacements) > 1 {
	// 	fmt.Println("---------------")
	// 	trigger.Print()
	//
	// 	fmt.Println("Replacement:")
	//
	// 	for inx, sli := range bestReplacements {
	//
	// 		fmt.Printf("inx => %d\n", inx)
	// 		sli.Print()
	// 	}
	//
	// 	fmt.Println("---------------")
	// }

	return success, 0, bestReplacements
}

func (slicer *Slicer) tryExpandShrink(queue *CoordinateQueue) {

	xy := queue.Pop()

	bestGain := 0
	var newSlice *pizza.Slice
	var sliceOverlaps []*pizza.Slice
	var sliceReplacements []*pizza.Slice
	// var leftovers []pizza.Coordinate

	for _, sliceCandidate := range slicer.SliceCache[ *xy ] {

		overlaps := slicer.overlapSlices(sliceCandidate)
		newSlices := make([]*pizza.Slice, 0)
		// newLeftovers := make([]pizza.Coordinate, 0)

		lost := 0
		replacementOk := true

		for _, shrinkSlice := range overlaps {

			status, sum, shrinked := slicer.shrinkSlice(sliceCandidate, shrinkSlice)

			if status == failed {
				replacementOk = false
				break
			}

			if status == eaten {
				continue
			}

			lost += shrinkSlice.Size() - sum
			newSlices = append(newSlices, shrinked...)

			// for _, leftXY := range newSlice.Complement(shrinkSlice) {
			// 	newLeftovers = append(newLeftovers, leftXY)
			// }
		}

		if !replacementOk {
			break
		}

		gain := sliceCandidate.Size() - lost

		if gain > bestGain {
			bestGain = gain
			newSlice = sliceCandidate
			sliceOverlaps = overlaps
			sliceReplacements = newSlices
			// leftovers = newLeftovers
		}
	}

	if newSlice == nil {
		return
	}

	for _, slice := range sliceOverlaps {
		slicer.Pizza.RemoveSlice(slice)
	}

	for _, slice := range sliceReplacements {
		slicer.Pizza.AddSlice(slice)
	}

	splitParts := slicer.splitSlice(newSlice)

	for _, slice := range splitParts {
		slicer.Pizza.AddSlice(slice)
	}
}

func (slicer *Slicer) ExpandThroughShrink() {

	fmt.Println("Expand through move")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		queue.Push(xy)
	}

	for queue.HasItems() {
		fmt.Printf("Move queue --> %-7d\r", len(queue.data) - 1)
		slicer.tryExpandShrink(queue)
	}

	fmt.Printf("Move queue --> done\n")
}
