package slicer

import (
	"../pizza"
	"fmt"
	"math/rand"
	"time"
)

func (slicer *Slicer) findRandom(xy pizza.Coordinate) *pizza.Slice {

	slices := slicer.SliceCache[ xy ]
	saveSlices := make([]*pizza.Slice, 0)

	for _, sli := range slices {

		if !slicer.overlap(sli) {
			saveSlices = append(saveSlices, sli)
		}
	}

	if len(saveSlices) <= 0 {
		return nil
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	inx := random.Intn(len(saveSlices))

	return saveSlices[ inx ]
}

func (slicer *Slicer) ExpandRandom() {

	fmt.Println("Expand random...")

	for _, xy := range slicer.Pizza.Traversal() {

		if slicer.Pizza.Cells[ xy ].Slice != nil {
			continue
		}

		slice := slicer.findRandom(xy)

		if slice == nil {
			continue
		}

		slicer.AddSlice(slice)
	}
}
