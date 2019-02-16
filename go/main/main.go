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

	// 100.00%
	// inputPath := "../../input/b_small.in"

	// 98.27%
	// inputPath := "../../input/c_medium.in"

	// 90.75%
	inputPath := "../../input/d_big.in"

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

	// pizz.CreateSubmission("submission-big.txt")
	// pizz.PrintSlices(false)
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
	// over.PrintVector()

	elapsed := time.Since(start)
	fmt.Printf("Done: %s\n", elapsed)
}
