package slicer

import (
	"../pizza"
)

func SearchSlices(pizza *pizza.Pizza) {

	slicer := Slicer{Pizza: pizza}

	slicer.Init()
	// slicer.FindBiggestParts()
	// slicer.FindSmallestParts()
	slicer.ExpandThroughDestruction()

	pizza.PrintSlices()
	pizza.PrintScore()

	// pizza.PrintSlicesToFile("xxx.txt")

	// start = &pizza.PizzaPart{
	// 	Pizza: pizz,
	// 	VectorR: pizza.Vector{Start: 13, End: 24},
	// 	VectorC: pizza.Vector{Start: 32, End: 47},
	// 	SliceCache: []*pizza.Slice{},
	// }
}
