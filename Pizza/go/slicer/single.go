package slicer

import "fmt"
import "../pizza"

func (slicer *Slicer) findSpecialCells() (noSlices []pizza.Coordinate, singles []pizza.Coordinate) {
	noSlicesCoords := make([]pizza.Coordinate, 0)
	oneSlicesPoss := make([]pizza.Coordinate, 0)

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		slices := slicer.SliceCache[ xy ]

		if len(slices) == 0 {
			noSlicesCoords = append(noSlicesCoords, xy)
			continue
		}

		if len(slices) > 1 {
			continue
		}

		oneSlicesPoss = append(oneSlicesPoss, xy)
	}

	return noSlicesCoords, oneSlicesPoss
}

func (slicer *Slicer) FindSingles() {

	fmt.Println("Find singles")

	noSlicesCoords, oneSlicesPoss := slicer.findSpecialCells()

	fmt.Printf("Not containable cells: %d\n", len(noSlicesCoords))
	fmt.Printf("Single containable cells: %d\n", len(oneSlicesPoss))

	slices := make([]*pizza.Slice, 0)

	for _, xy := range oneSlicesPoss {

		slice := slicer.SliceCache[ xy ][ 0 ]

		ok := true

		for inx, sli := range slices {

			if !slice.Overlap(sli) {
				continue
			}

			ok = false

			if sli.Size() < slice.Size() {
				slices[ inx ] = slice
			}
		}

		if ok {
			slices = append(slices, slice)
		}
	}

	for _, sli := range slices {
		slicer.Pizza.AddSlice(sli)
	}
}
