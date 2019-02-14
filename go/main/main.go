package main

import (
	"../pizza"
	"../slicer"
)

func main() {

	// 100.00%
	// inputPath := "../../input/a_example.in"

	// 100.00%
	// inputPath := "../../input/b_small.in"

	// 98.53%
	// inputPath := "../../input/c_medium.in"

	// 89.55%
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

	// pizz.PrintSlices(true)
	pizz.PrintScore()

	pizz.PrintSlicesToFile(true, "xxx.txt")

	//
	// over := pizza.Slice{
	// 	Pizza: &pizz,
	// 	Row: pizza.Vector{Start: 0, End: 1},
	// 	Column: pizza.Vector{Start: 2, End: 2},
	//
	// }
	// over.PrintVector()
}
