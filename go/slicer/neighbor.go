package slicer

import (
	"../pizza"
	"../simple"
	"fmt"
)

type Neighbor struct {
	Slices Slices
	Score  int
}

func (slicer *Slicer) calcNeighborFactor(slice *pizza.Slice) int {

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

	// return float32(factor) / float32(slice.Size())
	return factor - slice.Size()
}

func (slicer *Slicer) findBestNeighbor(xy pizza.Coordinate) *Neighbor {

	slices := slicer.SliceCache[ xy ]

	// factor := float32(0)
	factor := 0
	var slice *pizza.Slice

	for _, sli := range slices {

		if slicer.overlap(sli) {
			continue
		}

		neighborFactor := slicer.calcNeighborFactor(sli)

		if slice == nil || factor < neighborFactor {
			factor = neighborFactor
			slice = sli
		}
	}

	if slice == nil {
		return nil
	}

	return &Neighbor{Slices: slicer.splitSlice(slice), Score: factor}
}

func (slicer *Slicer) ExpandThroughNeighbors() {

	fmt.Println("Find best neighbors...")

	for _, xy := range slicer.Pizza.Traversal() {

		if slicer.Pizza.HasSliceAt(xy) {
			continue
		}

		neighbor := slicer.findBestNeighbor(xy)

		if neighbor != nil {
			slicer.AddSlices(neighbor.Slices)
		}
	}
}

func (slicer *Slicer) findBestNeighborCandidate(candidates map[pizza.Coordinate] *Neighbor) *Neighbor {

	score := 0
	var neighbor *Neighbor

	for _, candidate := range candidates {

		if candidate == nil {
			continue
		}

		if neighbor == nil || candidate.Score > score {
			score = candidate.Score
			neighbor = candidate
		}
	}

	return neighbor
}

func (slicer *Slicer) fixOverlapNeighbors(queue map[ pizza.Coordinate ] *Neighbor, bestSlice *pizza.Slice) {

	rowStart := simple.Max(slicer.Pizza.Row.Start, bestSlice.Row.Start - slicer.Pizza.MaxCells)
	rowEnd   := simple.Min(slicer.Pizza.Row.End,   bestSlice.Row.End   + slicer.Pizza.MaxCells)
	row := pizza.Vector{Start: rowStart, End: rowEnd}

	colStart := simple.Max(slicer.Pizza.Column.Start, bestSlice.Column.Start - slicer.Pizza.MaxCells)
	colEnd   := simple.Min(slicer.Pizza.Column.End,   bestSlice.Column.End   + slicer.Pizza.MaxCells)
	col := pizza.Vector{Start: colStart, End: colEnd}

	pseudoSlice := pizza.Slice{Row: row, Column: col}

	for _, xy := range pseudoSlice.Traversal() {

		if slicer.Pizza.HasSliceAt(xy) {
			delete(queue, xy)
			continue
		}

		if _, ok := queue[ xy ]; ok {
			best := slicer.findBestNeighbor(xy)

			if best == nil {
				delete(queue, xy)
			} else {
				queue[ xy ] = best
			}
		}
	}
}

func (slicer *Slicer) ExpandThroughNeighborsIntelligent() {

	fmt.Println("Find best neighbors...")

	// queue := make([]pizza.Coordinate, 1)
	// queue[ 0 ] = pizza.Coordinate{Row: slicer.Pizza.Row.Start, Column: slicer.Pizza.Column.Start}

	queue := make(map[ pizza.Coordinate ] *Neighbor)

	// startXY := pizza.Coordinate{Row: slicer.Pizza.Row.End, Column: slicer.Pizza.Column.End}
	// queue[ startXY ] = slicer.findBestNeighbor(startXY)

	for _, iny := range slicer.Pizza.Row.Range() {

		xy1 := pizza.Coordinate{Row: iny, Column: slicer.Pizza.Column.Start}
		xy2 := pizza.Coordinate{Row: iny, Column: slicer.Pizza.Column.End}

		queue[ xy1 ] = slicer.findBestNeighbor(xy1)
		queue[ xy2 ] = slicer.findBestNeighbor(xy2)
	}

	for _, inx := range slicer.Pizza.Column.Range() {

		xy1 := pizza.Coordinate{Row: slicer.Pizza.Row.Start, Column: inx}
		xy2 := pizza.Coordinate{Row: slicer.Pizza.Row.End, Column: inx}

		queue[ xy1 ] = slicer.findBestNeighbor(xy1)
		queue[ xy2 ] = slicer.findBestNeighbor(xy2)
	}

	covered := 0

	for len(queue) > 0 {

		best := slicer.findBestNeighborCandidate(queue)

		if best == nil {
			break
		}

		bestSlice := best.Slices
		slicer.AddSlices(bestSlice)

		for _, sli := range bestSlice {
			for _, xy := range sli.TraversalWithBorder() {

				if !slicer.Pizza.ContainsCoordinate(xy) {
					continue
				}

				if slicer.Pizza.HasSliceAt(xy) {
					continue
				}

				queue[ xy ] = slicer.findBestNeighbor(xy)
			}

			slicer.fixOverlapNeighbors(queue, sli)
			covered += sli.Size()
		}

		fmt.Printf("covered=%d queue=%-6d\r", covered, len(queue),)
	}

	fmt.Println()
}
