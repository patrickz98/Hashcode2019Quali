package main

import (
	"../pizza"
	"../slicer"
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// 100.00%
	// inputPath := "../../input/a_example.in"
	// submissionPath := "../../submissions/a_example.out"

	// 100.00%
	// inputPath := "../../input/b_small.in"
	// submissionPath := "../../submissions/b_small.out"

	// 98.43%
	inputPath := "../../input/c_medium.in"
	submissionPath := "../../submissions/c_medium.out"

	// 90.83%
	// inputPath := "../../input/d_big.in"
	// submissionPath := "../../submissions/d_big.out"

	pizz := pizza.NewPizza(inputPath)

	// pizz = pizza.Pizza{
	// 	Ingredients: pizz.Ingredients,
	// 	MaxCells:    pizz.MaxCells,
	// 	Cells:       pizz.Cells,
	// 	Row: pizza.Vector{Start: 0, End: 99},
	// 	Column: pizza.Vector{Start: 0, End: 99},
	// }

	pizz.PrintParams()

	slicer.SearchSlices(&pizz)
	// pizz.CheckErrors()

	// pizz.PrintSlices(false)
	pizz.CreateSubmission(submissionPath)
	pizz.PrintScore()

	// pizz.PrintSlicesToFile(true, "xxx.txt")
	// pizz.PrintSlicesToFile(false, "yyy.txt")

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
