package slicer

import (
	"../pizza"
	"fmt"
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
	// slicer.TryDestuctShink()
	// slicer.ExpandBalancedIntelligent()

	// med: 98.60%
	// slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()

	// med: 98.58%
	// slicer.ExpandThroughNeighbors()
	// slicer.TryDestuctShink()

	// med: 98.83%
	// slicer.ExpandThroughNeighbors()
	// slicer.TryDestuctShink()
	// slicer.MoveSlices()
	// slicer.TryDestuctShink()

	// med: 99.07%
	// slicer.ExpandThroughNeighbors()
	// slicer.TryDestuctShink()
	// slicer.ShakeSlices()
	// slicer.TryDestuctShink()

	// slicer.MoveSlices()
	// slicer.TryDestuctShink()

	// big: 86.31% med: 95.03%
	// slicer.ExpandBalanced()
	// slicer.TryDestuctShink()

	// big: 90.83% med: 98.62%
	// slicer.ExpandThroughNeighbors()
	// slicer.TryDestuctShink()

	// big: 90.97% med: 98.80%
	// slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()
	// slicer.TryDestuctShink()

	// big: 92.42% mid: 99.19%
	slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()
	// slicer.TryDestuctShink()

	bestCover, _ := piz.Score()

	for {
		// slicer.ExpandThroughDestruction()
		// slicer.ExpandThroughShrink()
		slicer.TryDestuctShink()
		slicer.ShakeSlices()

		cover, _ := piz.Score()

		fmt.Printf("############# cover=%d\n", cover)

		if bestCover == cover {
			break
		}

		bestCover = cover
	}
}
