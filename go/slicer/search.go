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
	// slicer.ExpandBalanced()
	// slicer.TryAll()

	// big: 86.31% med: 95.03%
	slicer.ExpandBalanced()
	slicer.TryAll()

	// big: 90.82% med: 98.54%
	// slicer.ExpandThroughNeighbors()
	// slicer.TryAll()

	// big: 91.13% med: 98.80%
	// slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()
	// slicer.TryAll()

	// big: 92.39%
	// slicer.FindSingles()
	// slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()
	//
	// bestCover, _ := piz.Score()
	//
	// for {
	// 	fmt.Printf("############# cover=%d\n", bestCover)
	//
	// 	slicer.TryAll()
	// 	slicer.ChangeSlices()
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
