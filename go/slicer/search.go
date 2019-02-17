package slicer

import (
	"../pizza"
	"fmt"
)

func SearchSlices(pizza *pizza.Pizza) {

	slicer := Slicer{Pizza: pizza}
	slicer.Init()

	// slicer.FindBiggestParts()

	// slicer.FindSingles()
	slicer.FindSmallestParts()

	slicer.ExpandThroughDestruction()
	slicer.ExpandThroughShrink()

	covered := pizza.SliceCount()

	for {
		fmt.Printf("covered=%d\n", covered)

		slicer.ExpandThroughMove()

		slicer.FindBiggestParts()
		slicer.ExpandThroughDestruction()
		slicer.ExpandThroughShrink()

		covered2 := pizza.SliceCount()

		if covered == covered2 {
			break
		} else {
			covered = covered2
		}
	}
}
