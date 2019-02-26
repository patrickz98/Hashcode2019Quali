
package cornerSetter

import (
	"../pizza"
	"../slicer"
)

func SearchSlices(piz *pizza.Pizza) {

	slicer := slicer.Slicer{Pizza: piz}
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

	// big: 90.82% med: 98.54%
	// slicer.ExpandThroughNeighbors()
	// slicer.TryAll()

	// big: 91.13% med: 98.80%
	// slicer.ExpandThroughNeighbors()
	// slicer.ExpandThroughDestruction()
	// slicer.ExpandThroughShrink()
	// slicer.TryAll()

	slicer.FindSingles()
	slicer.ExpandThroughNeighbors()
	slicer.ExpandThroughDestruction()
	slicer.ExpandThroughShrink()

	bestCover, _ := piz.Score()

	for {
		fmt.Printf("############# cover=%d\n", bestCover)

		slicer.TryAll()
		slicer.ChangeSlices()

		cover, _ := piz.Score()

		if bestCover == cover {
			break
		}

		bestCover = cover
	}
}