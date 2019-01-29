package main

import (
	"../pizza"
	"../slicer"
	"fmt"
)

func main() {

	fmt.Println("Start")

	inputPath := "../../input/a_example.in"
	// inputPath := "../../input/b_small.in"
	// inputPath := "../../input/c_medium.in"
	// inputPath := "../../input/d_big.in"

	pizz := pizza.NewPizza(inputPath)


	// cuters := make([]pizza.Slice, 1)
	// cuters[ 0 ] = pizza.Slice{
	// 	Row: pizza.Vector{Start:0, End: 2},
	// 	Column: pizza.Vector{Start:0, End: 0},
	// }
	// cuters[ 0 ] = pizza.Slice{
	// 	Row: pizza.Vector{Start:0, End: 0},
	// 	Column: pizza.Vector{Start:0, End: 4},
	// }
	// cuters[ 0 ] = pizza.Slice{
	// 	Row: pizza.Vector{Start:1, End: 2},
	// 	Column: pizza.Vector{Start:0, End: 1},
	// }
	// cuters[ 0 ] = pizza.Slice{
	// 	Row: pizza.Vector{Start:0, End: 2},
	// 	Column: pizza.Vector{Start:0, End: 0},
	// }
	// cuters[ 1 ] = pizza.Slice{
	// 	Row: pizza.Vector{Start:0, End: 0},
	// 	Column: pizza.Vector{Start:1, End: 3},
	// }

	// pizz.Slices = cuters

	pizz.PrintParams()
	// pizz.PrintSlices()

	slicer.FindSlice(&pizz)

	// pizz.PrintPizzaCells()
	pizz.PrintSlices()
	pizz.PrintScore()
}
