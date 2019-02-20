package slicer

import (
	"../pizza"
	"fmt"
	"math/rand"
	"time"
)

func (slicer *Slicer) findRandom(xy pizza.Coordinate) {

	if slicer.Pizza.Cells[ xy ].Slice != nil {
		return
	}

	slices := slicer.SliceCache[ xy ]
	saveSlices := make([]*pizza.Slice, 0)

	for _, sli := range slices {

		if !slicer.overlap(sli) {
			saveSlices = append(saveSlices, sli)
		}
	}

	if len(saveSlices) <= 0 {
		return
	}

	// fmt.Printf("len(saveSlices)=%d\n", len(saveSlices))

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	inx := random.Intn(len(saveSlices))

	// fmt.Printf("inx=%d\n", inx)

	slicer.Pizza.AddSlice(saveSlices[ inx ])
}

func (slicer *Slicer) ExpandRandom() {

	fmt.Println("Expand random...")

	for _, xy := range slicer.Pizza.Traversal() {
		slicer.findRandom(xy)
	}
}
