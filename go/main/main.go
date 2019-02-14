package main

import (
	"../pizza"
	"../slicer"
)

func main() {

	// 100.00%
	// inputPath := "../../input/a_example.in"

	// 97.62%
	inputPath := "../../input/b_small.in"

	// 98.53%
	// inputPath := "../../input/c_medium.in"

	// 89.55%
	// inputPath := "../../input/d_big.in"

	pizz := pizza.NewPizza(inputPath)
	pizz.PrintParams()
	// pizz.PrintPizza()

	// ori := pizza.Slice{
	// 	Pizza: &pizz,
	// 	Row: pizza.Vector{Start: 0, End: 2},
	// 	Column: pizza.Vector{Start: 0, End: 1},
	//
	// }
	// ori.PrintVector()
	//
	// over := pizza.Slice{
	// 	Pizza: &pizz,
	// 	Row: pizza.Vector{Start: 0, End: 1},
	// 	Column: pizza.Vector{Start: 2, End: 2},
	//
	// }
	// over.PrintVector()
	//
	// fmt.Println(ori.Overlap(&over))

	slicer.SearchSlices(&pizz)
	// slicer.FindSlice(part)

	// pizz.PrintSlices()

	// parts := pizz.Cut()
	// printParts(parts)

	// part := parts[ 2 ]
	// part.PrintPizza()
	// slicer.FindSlice(part)
	// part.PrintSlices()

	// slicer.FindSlice(&pizz)

	// pizz.PrintPizzaCells()
	// pizz.PrintSlices()
	// pizz.PrintSlicesPlain()
	// pizz.PrintScore()
}
