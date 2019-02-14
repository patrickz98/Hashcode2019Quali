package slicer

import (
	"../pizza"
	"fmt"
)

type Slicer struct {
	Pizza      *pizza.Pizza
	SliceCache map[pizza.Coordinate] []*pizza.Slice
}

func (slicer *Slicer) Init() {
	slicer.buildSlicesCache()
}

func (slicer *Slicer) overlap(slice *pizza.Slice) bool {

	// TODO: Optimise
	for _, xy := range slice.Traversal() {

		cell := slicer.Pizza.Cells[ xy ]
		if cell.Slice != nil {
			return true
		}
	}

	return false
}

func (slicer *Slicer) overlapSlices(slice *pizza.Slice) []*pizza.Slice {

	overlap := make([]*pizza.Slice, 0)

	for _, xy := range slice.Traversal() {

		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil && !contains(overlap, cell.Slice) {
			overlap = append(overlap, cell.Slice)
		}
	}

	return overlap
}

func (slicer *Slicer) FindSmallestParts() {

	size := slicer.Pizza.Size()

	for count, xy := range slicer.Pizza.Traversal() {

		slices := slicer.SliceCache[ xy ]

		var smallest *pizza.Slice

		for _, slice := range slices {

			if slice == nil {
				continue
			}

			if slicer.overlap(slice) {
				continue
			}

			// slic.PrintInfo()

			if (smallest == nil) || (smallest.Size() > slice.Size()) {
				smallest = slice
			}
		}

		if smallest != nil {
			slicer.Pizza.AddSlice(smallest)
		}

		fmt.Printf("Find smalest slices %d/%d\r", size, count + 1)
	}

	fmt.Println()
}

func (slicer *Slicer) FindBiggestParts() {

	size := slicer.Pizza.Size()

	for count, xy := range slicer.Pizza.Traversal() {

		slices := slicer.SliceCache[ xy ]

		var bigggest *pizza.Slice

		for _, slice := range slices {

			if slice == nil {
				continue
			}

			if slicer.overlap(slice) {
				continue
			}

			// slic.PrintInfo()

			if (bigggest == nil) || (bigggest.Size() < slice.Size()) {
				bigggest = slice
			}
		}

		if bigggest != nil {
			slicer.Pizza.AddSlice(bigggest)
		}

		fmt.Printf("Find biggest slices %d/%d\r", size, count + 1)
	}

	fmt.Println()
}