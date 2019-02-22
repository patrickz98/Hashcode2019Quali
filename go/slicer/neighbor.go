package slicer

import (
	"../pizza"
	"fmt"
)

type Neighbor struct {
	Slice *pizza.Slice
	Score float32
}

func (slicer *Slicer) calcNeighborFactor(slice *pizza.Slice) float32 {

	rowStart := slice.Row.Start - 1
	rowEnd := slice.Row.End + 1

	colStart := slice.Column.Start - 1
	colEnd := slice.Column.End + 1

	factor := 0

	for iny := rowStart; iny <= rowEnd; iny++ {
		for inx := colStart; inx <= colEnd; inx++ {

			xy := pizza.Coordinate{Row: iny, Column: inx}

			if slice.ContainsCoordinate(xy) {
				continue
			}

			if !slicer.Pizza.ContainsCoordinate(xy) {
				factor++
				continue
			}

			if slicer.Pizza.HasSliceAt(xy) {
				factor++
			}
		}
	}

	return float32(factor) / float32(slice.Size())
}

func (slicer *Slicer) findBestNeighbor(xy pizza.Coordinate) *Neighbor {

	slices := slicer.SliceCache[ xy ]

	factor := float32(0)
	var slice *pizza.Slice

	for _, sli := range slices {

		if slicer.overlap(sli) {
			continue
		}

		neighborFactor := slicer.calcNeighborFactor(sli)

		if factor < neighborFactor {
			factor = neighborFactor
			slice = sli
		}
	}

	if slice == nil {
		return nil
	}

	return &Neighbor{Slice: slice, Score: factor}
}

func (slicer *Slicer) ExpandThroughNeighbors() {

	fmt.Println("Expand edge...")

	for _, xy := range slicer.Pizza.Traversal() {

		if !slicer.Pizza.HasSliceAt(xy) {

			neighbor := slicer.findBestNeighbor(xy)

			if neighbor != nil {
				slicer.Pizza.AddSlice(neighbor.Slice)
			}
		}
	}
}

func (slicer *Slicer) findBestNeighborCandidate(candidates map[pizza.Coordinate] *Neighbor) *Neighbor {

	score := float32(0)
	var neighbor *Neighbor

	for _, candidate := range candidates {

		if candidate == nil {
			continue
		}

		if candidate.Score > score {
			score = candidate.Score
			neighbor = candidate
		}
	}

	return neighbor
}

func (slicer *Slicer) ExpandThroughNeighborsIntelligent() {

	fmt.Println("Expand edge...")

	// queue := InitCoordinateQueue()
	// queue.Push()

	queue := make([]pizza.Coordinate, 1)
	queue[ 0 ] = pizza.Coordinate{Row: slicer.Pizza.Row.Start, Column: slicer.Pizza.Column.Start}

	coverd := 0

	for len(queue) > 0 {

		scores := make(map[pizza.Coordinate] *Neighbor)

		for _, xy := range queue {

			if !slicer.Pizza.HasSliceAt(xy) {
				scores[ xy ] = slicer.findBestNeighbor(xy)
			} else {
				scores[ xy ] = nil
			}
		}

		best := slicer.findBestNeighborCandidate(scores)

		if best == nil {
			break
		}

		bestSlice := best.Slice

		slicer.Pizza.AddSlice(bestSlice)

		tmp := make([]pizza.Coordinate, 0)

		rowStart := bestSlice.Row.Start - 1
		rowEnd := bestSlice.Row.End + 1

		colStart := bestSlice.Column.Start - 1
		colEnd := bestSlice.Column.End + 1

		for iny := rowStart; iny <= rowEnd; iny++ {
			for inx := colStart; inx <= colEnd; inx++ {

				xy := pizza.Coordinate{Row: iny, Column: inx}

				if bestSlice.ContainsCoordinate(xy) {
					continue
				}

				if !slicer.Pizza.ContainsCoordinate(xy) {
					continue
				}

				if !slicer.Pizza.HasSliceAt(xy) {
					tmp = append(tmp, xy)
				}
			}
		}

		for xy, neighbor := range scores {

			if best == neighbor {
				continue
			}

			if slicer.Pizza.HasSliceAt(xy) {
				continue
			}

			if neighbor != nil {
				tmp = append(tmp, xy)
			}
		}

		queue = tmp

		coverd += bestSlice.Size()

		bestSlice.PrintVector()
		fmt.Printf("queue=%d coverd=%d\n", len(queue), coverd)

		// for _, xy := range queue {
			// fmt.Printf("(%d, %d)\n", xy.Row, xy.Column)
		// }
	}
}
