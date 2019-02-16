package slicer

import (
	"../pizza"
)

func SearchSlices(pizza *pizza.Pizza) {

	slicer := Slicer{Pizza: pizza}

	slicer.Init()

	slicer.FindBiggestParts()

	// slicer.FindSingles()
	// slicer.FindSmallestParts()

	slicer.ExpandThroughDestruction()
	slicer.ExpandThroughShrink()

	// covered := pizza.Covered()
	//
	// for {
	// 	slicer.ExpandThroughMove()
	//
	// 	slicer.FindBiggestParts()
	// 	slicer.ExpandThroughDestruction()
	// 	slicer.ExpandThroughShrink()
	//
	// 	covered2 := pizza.Covered()
	//
	// 	if covered == covered2 {
	// 		break
	// 	} else {
	// 		covered = covered2
	// 	}
	// }
}
