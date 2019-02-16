package slicer

import "fmt"
import "../pizza"

func (slicer *Slicer) FindSingles() {

	fmt.Println("Find singles")

	noSlicesCoords := make([]pizza.Coordinate, 0)
	oneSlicesPoss := make([]*pizza.Slice, 0)

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

		if len(slices) == 1 {

			if ! contains(oneSlicesPoss, slices[ 0 ]) {
				oneSlicesPoss = append(oneSlicesPoss, slices[ 0 ])
			}

			continue
		}

	}

	fmt.Printf("Not containable cells: %d\n", len(noSlicesCoords))
	fmt.Printf("Single containable cells: %d\n", len(oneSlicesPoss))

	for _, sli := range oneSlicesPoss {
		slicer.Pizza.AddSlice(sli)
	}
}
