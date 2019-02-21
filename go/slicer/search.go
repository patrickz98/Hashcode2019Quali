package slicer

import (
	"../pizza"
	"fmt"
)

func SearchSlices(piz *pizza.Pizza, params []SlicerParams) {

	slicer := Slicer{Pizza: piz}
	slicer.Init()

	// slicer.ExpandRandom()

	scorePercent := make([]float32, 0)

	for _, param := range params {
		fmt.Println("\nNow trying: " + param.Name)

		piz.RemoveAllSlice()
		slicer.Params = param
		slicer.setCorners()

		//slicer.FindBiggestParts()
		//slicer.FindSingles()

		//slicer.ExpandThroughEdge()
		slicer.FindSmallestParts()
		slicer.ExpandThroughDestruction()
		slicer.ExpandThroughShrink()

		_, addToScore := piz.Score()
		scorePercent = append(scorePercent, addToScore*100)
		piz.PrintScore()
	}

	fmt.Println()

	for i, param := range params {
		fmt.Printf(param.Name+": %f\n", scorePercent[i])
	}

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
