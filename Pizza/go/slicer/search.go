package slicer

import (
	"../pizza"
	"fmt"
)

func brute1(slicer *Slicer) {

	slicer.ExpandThroughNeighbors()

	bestCover, _ := slicer.Pizza.Score()

	for {
		slicer.TryDestuctShink()
		slicer.ShakeSlices()

		cover, _ := slicer.Pizza.Score()

		fmt.Printf("############# cover=%d\n", cover)

		if bestCover == cover {
			break
		}

		bestCover = cover
	}
}

func brute2(slicer *Slicer) {

	slicer.ExpandThroughNeighbors()
	slicer.TryDestuctShink()

	bestCover, _ := slicer.Pizza.Score()

	for {
		//slicer.TryDestuctShink()
		slicer.DestructShrinkHoles()
		slicer.ShakeSlices()

		cover, _ := slicer.Pizza.Score()

		fmt.Printf("############# cover=%d\n", cover)

		if bestCover == cover {
			break
		}

		bestCover = cover
	}
}

func brute3(slicer *Slicer) {

	slicer.ExpandThroughNeighbors()
	slicer.TryDestuctShink()

	bestCover, _ := slicer.Pizza.Score()

	for {
		slicer.TryDestuctShink()
		slicer.ShakeHoles()

		cover, _ := slicer.Pizza.Score()

		fmt.Printf("############# cover=%d\n", cover)

		if bestCover == cover {
			break
		}

		bestCover = cover
	}
}

func brute4(slicer *Slicer) {

	slicer.ExpandThroughNeighbors()
	slicer.TryDestuctShink()

	bestCover, _ := slicer.Pizza.Score()

	for {

		slicer.ShakeHolesFill()

		cover, _ := slicer.Pizza.Score()

		fmt.Printf("############################### cover=%d\n", cover)

		if bestCover == cover {
			break
		}

		bestCover = cover
	}
}

func SearchSlices(piz *pizza.Pizza) {

	slicer := &Slicer{Pizza: piz}
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

	//slicer.ExpandThroughNeighbors()
	////slicer.TryDestuctShink()
	//slicer.DestructShrinkHoles()
	//slicer.TryDestuctShink()

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

	// big: 92.42% med: 99.19%
	//brute1(slicer)

	// big: 00.00% med: 99.18%
	//brute2(slicer)

	// big: 93.80% med: 99.25%
	//brute3(slicer)

	brute4(slicer)
}
