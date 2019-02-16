package slicer

import (
	"../pizza"
)

func contains(slices []*pizza.Slice, slice *pizza.Slice) bool {

	for _, sli := range slices {
		if sli == slice {
			return true
		}
	}

	return false
}

func (slicer Slicer) slicesInSlice(slice *pizza.Slice) []*pizza.Slice {

	parts := make([]*pizza.Slice, 0)

	for _, xy := range slice.Traversal() {
		for _, sli := range slicer.SliceCache[ xy ] {

			if contains(parts, sli) {
				continue
			}

			if !slice.Contains(sli) {
				continue
			}

			parts = append(parts, sli)
		}
	}

	return parts
}

// Split existing slice in small peaces.
func (slicer Slicer) splitSlice(slice *pizza.Slice) []*pizza.Slice {

	possibleParts := make([]*pizza.Slice, 0)
	possibleParts = append(possibleParts, slice)

	ingredientsCount := slicer.Pizza.Ingredients

	for _, sli := range slicer.slicesInSlice(slice) {

		if slice == sli {
			continue
		}

		leftOver := slice.Size() - sli.Size()

		if leftOver < (ingredientsCount * 2) {
			continue
		}

		possibleParts = append(possibleParts, sli)
	}

	powerSet := make([][]*pizza.Slice, 0)

	for _, slicePart := range possibleParts {

		tmp := make([]*pizza.Slice, 0)
		tmp = append(tmp, slicePart)

		for _, rr := range powerSet {
			tmp2 := append(tmp, rr...)
			powerSet = append(powerSet, tmp2)
		}

		powerSet = append(powerSet, tmp)
	}

	newSlices := make([]*pizza.Slice, 0)

	for _, set := range powerSet {

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

		if sum != slice.Size() {
			continue
		}

		if len(newSlices) < len(set) {
			newSlices = set
		}
	}

	if len(newSlices) == 0 {
		newSlices = append(newSlices, slice)
	}

	return newSlices
}
