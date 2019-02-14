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
	slicer.ExpandThroughMove()
}
