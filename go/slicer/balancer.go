package slicer

import "../pizza"

func (slicer *Slicer) findBalanced(xy pizza.Coordinate) *pizza.Slice {

	if slicer.Pizza.HasSliceAt(xy) {
		return nil
	}

	slices := slicer.SliceCache[ xy ]

	for _, sli := range slices {

		balance := sli.IngredientsBalance()

		if balance != 0 {
			continue
		}

		// if sli.Size() != slicer.Pizza.MaxCells {
		// 	continue
		// }

		if slicer.overlap(sli) {
			continue
		}

		return sli
	}

	return nil
}

func (slicer *Slicer) ExpandBalanced() {

	for _, xy := range slicer.Pizza.Traversal() {

		slice := slicer.findBalanced(xy)

		if slice == nil {
			continue
		}

		slicer.AddSlice(slice)
	}
}
