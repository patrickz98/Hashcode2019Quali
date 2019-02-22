package slicer

import (
	"../pizza"
)

func SearchSlices(piz *pizza.Pizza) {

	slicer := Slicer{Pizza: piz}
	slicer.Init()

	// slicer.ExpandRandom()
	// slicer.FindSingles()
	// slicer.ExpandThroughNeighborsIntelligent()
	// slicer.ExpandThroughNeighbors()
	// slicer.FindSmallestParts()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()
	// slicer.ExpandThroughCorners()
	// slicer.ExpandShot()
	// slicer.FindBiggestParts()

	// slicer.TryAll()

	slicer.ExpandThroughNeighbors()
	slicer.ExpandThroughDestruction()
	slicer.ExpandThroughShrink()
	slicer.TryAll()

	// bestCover, _ := piz.Score()
	//
	// for {
	// 	fmt.Printf("############# cover=%d\n", bestCover)
	//
	// 	// slicer.ExpandThroughDestruction()
	// 	// slicer.ExpandThroughShrink()
	// 	slicer.TryAll()
	// 	slicer.MoveSlices()
	//
	// 	cover, _ := piz.Score()
	//
	// 	if bestCover == cover {
	// 		break
	// 	}
	//
	// 	bestCover = cover
	// }
}
