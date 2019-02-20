package slicer

import (
	"../pizza"
)

func SearchSlices(piz *pizza.Pizza) {

	slicer := Slicer{Pizza: piz}
	slicer.Init()

	// slicer.ExpandRandom()

	// slicer.FindBiggestParts()
	// slicer.FindSingles()

	slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughNeighborsIntelligent()
	// slicer.FindSmallestParts()
	slicer.ExpandThroughDestruction()
	slicer.ExpandThroughShrink()

	// slicer.ExpandShot()
	// slicer.FindBiggestParts()

	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()

	// bestCover := 0
	// var slices []*pizza.Slice
	//
	// for inx := 0; inx < 10; inx++ {
	//
	// 	slicer.FindSingles()
	// 	slicer.ExpandRandom()
	// 	// slicer.ExpandThroughDestruction()
	// 	// slicer.ExpandThroughShrink()
	//
	// 	cover, _ := piz.Score()
	//
	// 	if bestCover < cover {
	// 		bestCover = cover
	// 		slices = piz.Slices()
	// 	}
	//
	// 	fmt.Printf("############# %d cover=%d\n", inx, cover)
	//
	// 	piz.RemoveAllSlice()
	// }
	//
	// fmt.Printf("############# best=%d\n", bestCover)
	//
	// for _, sli := range slices {
	// 	piz.AddSlice(sli)
	// }
}
