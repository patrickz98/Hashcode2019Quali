package main

import (
	"../pizza"
	"../slicer"
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// best: 100.00% --> 100.00%
	// inputPath := "../../input/a_example.in"
	// submissionPath := "../../submissions/a_example.out"

	// best: 100.00% --> 100.00%
	// inputPath := "../../input/b_small.in"
	// submissionPath := "../../submissions/b_small.out"

	// best: 99.33% --> 99.30%
	// inputPath := "../../input/c_medium.in"
	// submissionPath := "../../submissions/c_medium.out"

	// best: 93.06% --> 91.19%
	inputPath := "../../input/d_big.in"
	// submissionPath := "../../submissions/d_big.out"

	paramSlices := make([]slicer.SlicerParams, 0)

	paramSlices = append(paramSlices, slicer.SlicerParams{
		"Corner 0/3", 0, 3,
	})

	paramSlices = append(paramSlices, slicer.SlicerParams{
		"Corner 50/3", 50, 3,
	})

	paramSlices = append(paramSlices, slicer.SlicerParams{
		"Corner 100/3", 100, 3,
	})

	paramSlices = append(paramSlices, slicer.SlicerParams{
		"Corner 150/3", 150, 3,
	})

	paramSlices = append(paramSlices, slicer.SlicerParams{
		"Corner 500/3", 500, 3,
	})

	paramSlices = append(paramSlices, slicer.SlicerParams{
		"Corner 1000/3", 500, 3,
	})

	piz := pizza.NewPizza(inputPath)

	slicer.SearchSlices(&piz, paramSlices)

	//piz := pizza.NewPizza(inputPath)

	// piz = pizza.Pizza{
	// 	Ingredients: piz.Ingredients,
	// 	MaxCells:    piz.MaxCells,
	// 	Cells:       piz.Cells,
	// 	Row: pizza.Vector{Start: 0, End: 59},
	// 	Column: pizza.Vector{Start: 0, End: 79},
	// }

	// piz.PrintParams()

	//slicer.SearchSlices(&piz)
	// piz.CheckErrors()

	// piz.PrintSlices(false)
	// piz.CreateSubmission(submissionPath)
	//piz.PrintScore()

	// piz.PrintSlicesToFile(true, "yyy-marked.txt")
	// piz.PrintSlicesToFile(false, "yyy.txt")

	//
	// over := pizza.Slice{
	// 	Pizza: &pizz,
	// 	Row: pizza.Vector{Start: 0, End: 1},
	// 	Column: pizza.Vector{Start: 2, End: 2},
	//
	// }
	// over.VectorPrint()

	elapsed := time.Since(start)
	fmt.Printf("Done: %s\n", elapsed)
}
