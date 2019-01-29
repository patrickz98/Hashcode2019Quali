package main

import (
	"../pizza"
	"../slicer"
	"fmt"
)

var count = 0

func cut(piz *pizza.Pizza) {

	count++
	fmt.Printf("count=%d\n", count)

	if ! piz.CutPossible() {
		fmt.Println("Done cutting...")
		piz.PrintPizza()
		fmt.Println("")
		return
	}

	fmt.Println("Cutting...")
	piz.PrintPizza()
	fmt.Println("")

	parts := piz.Cut()

	for _, val := range parts {

		if val != nil {
			cut(val)

			// merge

			break
		}
	}
}

func main() {

	for inx := 0; inx < 2; inx++ {

	}

	fmt.Println("Start")

	inputPath := "/Users/patrick/Desktop/google/input/a_example.in"
	// inputPath := "/Users/patrick/Desktop/google/input/b_small.in"
	// inputPath := "/Users/patrick/Desktop/google/input/c_medium.in"
	// inputPath := "/Users/patrick/Desktop/google/input/d_big.in"

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

	finder := slicer.Finder{Pizza: &pizz}
	finder.FindSlice()

	// pizz.PrintPizzaCells()
	pizz.PrintSlices()
	pizz.PrintScore()
}
