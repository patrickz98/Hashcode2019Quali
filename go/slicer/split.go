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

	ingredientsCount := slicer.Pizza.Ingredients

	// slice.Print()

	for _, sli := range slicer.slicesInSlice(slice) {

		if slice == sli {
			continue
		}

		leftOver := slice.Size() - sli.Size()

		if leftOver < (ingredientsCount * 2) {
			continue
		}

		possibleParts = append(possibleParts, sli)

		// fmt.Println("Possible part:")
		// sli.Print()
	}

	// if true {
	// 	fmt.Println("Bye")
	// 	os.Exit(1)
	// }

	for _, slix := range possibleParts {
		for _, sliy := range possibleParts {

			if slix.Overlap(sliy) {
				continue
			}

			sum := slix.Size() + sliy.Size()

			if sum < slice.Size() {
				continue
			}

			newSlices := make([]*pizza.Slice, 0)
			newSlices = append(newSlices, slix)
			newSlices = append(newSlices, sliy)

			return newSlices
		}
	}

	parts := make([]*pizza.Slice, 0)
	parts = append(parts, slice)
	return parts
}
