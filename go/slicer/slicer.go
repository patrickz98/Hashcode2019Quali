package slicer

import (
	"../pizza"
	"fmt"
)

type Slices []*pizza.Slice

type Slicer struct {
	Pizza             *pizza.Pizza
	SliceCache        map[ pizza.Coordinate ] Slices
	TopLeftSliceCache map[ pizza.Coordinate ] Slices
}

func (slicer *Slicer) Init() {
	slicer.buildSlicesCache()
}

func (slicer Slicer) CalculateSize(slices Slices) int {

	size := 0

	for _, sil := range slices {
		size += sil.Size()
	}

	return size
}

func (slicer Slicer) CalculateGain(slices1 Slices, slices2 Slices) int {

	return slicer.CalculateSize(slices1) - slicer.CalculateSize(slices2)
}

func (slicer *Slicer) overlap(slice *pizza.Slice) bool {

	for _, xy := range slice.Traversal() {

		cell := slicer.Pizza.Cells[xy]

		if cell.Slice != nil {
			return true
		}
	}

	return false
}

func (slicer *Slicer) overlapSlices(slice *pizza.Slice) []*pizza.Slice {

	overlap := make([]*pizza.Slice, 0)

	for _, xy := range slice.Traversal() {

		cell := slicer.Pizza.Cells[xy]

		if cell.Slice != nil && !contains(overlap, cell.Slice) {
			overlap = append(overlap, cell.Slice)
		}
	}

	return overlap
}

func (slicer *Slicer) findSmallestAt(xy pizza.Coordinate) *pizza.Slice {

	var smallest *pizza.Slice

	slices := slicer.SliceCache[xy]

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

	return smallest
}

func (slicer *Slicer) FindSmallestParts() {

	size := slicer.Pizza.Size()

	for count, xy := range slicer.Pizza.Traversal() {

		fmt.Printf("Find smalest slices %d/%d\r", size, count+1)

		if slicer.Pizza.Cells[xy].Slice != nil {
			continue
		}

		smallest := slicer.findSmallestAt(xy)

		if smallest != nil {
			slicer.Pizza.AddSlice(smallest)
		}
	}

	fmt.Println()
}

func (slicer *Slicer) findBiggestAt(xy pizza.Coordinate) *pizza.Slice {

	slices := slicer.SliceCache[xy]

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

	return bigggest
}

func (slicer *Slicer) FindBiggestParts() {

	size := slicer.Pizza.Size()

	for count, xy := range slicer.Pizza.Traversal() {

		fmt.Printf("Find biggest slices %d/%d\r", size, count+1)

		if slicer.Pizza.Cells[xy].Slice != nil {
			continue
		}

		biggest := slicer.findBiggestAt(xy)

		if biggest != nil {
			slicer.Pizza.AddSlice(biggest)
		}
	}

	fmt.Println()
}

func (slicer *Slicer) AddSlice(slice *pizza.Slice) {

	slicer.Pizza.AddSlice(slice)
}

func (slicer *Slicer) AddSlices(slices Slices) {

	for _, slice := range slices {
		slicer.Pizza.AddSlice(slice)
	}
}

func (slicer *Slicer) RemoveSlice(slice *pizza.Slice) {

	slicer.Pizza.RemoveSlice(slice)
}

func (slicer *Slicer) RemoveSlices(slices Slices) {

	for _, slice := range slices {
		slicer.Pizza.RemoveSlice(slice)
	}
}

func (slicer *Slicer) leftover(slice *pizza.Slice, overlap *pizza.Slice) []pizza.Coordinate {

	return overlap.Complement(slice)
}

func (slicer *Slicer) leftovers(slices Slices, overlaps Slices) []pizza.Coordinate {

	coords := make([]pizza.Coordinate, 0)

	// fmt.Println("-------- slices --------")
	//
	// for _, tmp := range slices {
	// 	tmp.PrintVector()
	// 	tmp.Print()
	// }
	//
	// fmt.Println("------- overlaps -------")
	//
	// for _, tmp := range overlaps {
	// 	tmp.PrintVector()
	// 	tmp.Print()
	// }

	for _, overlap := range overlaps {

		for _, sli := range slices {

			coords = append(coords, overlap.Complement(sli)...)
		}
	}

	// fmt.Println("------- coords -------")
	//
	// for _, coord := range coords {
	// 	fmt.Println(coord.Stringify())
	// }
	//
	// simple.Exit()

	return coords
}
