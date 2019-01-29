package main

import (
	"../pizza"
	"../slicer"
	"fmt"
)

func main() {

	fmt.Println("Start")

	// vec := pizza.Vector{Start: 2, End: 4}
	//
	// for _, inx := range vec.Range() {
	// 	fmt.Printf("inx=%d\n", inx)
	// }

	// inputPath := "../../input/a_example.in"
	inputPath := "../../input/b_small.in"
	// inputPath := "../../input/c_medium.in"
	// inputPath := "../../input/d_big.in"

	pizz := pizza.NewPizza(inputPath)

	// pizz.Slices = cuters

	pizz.PrintParams()
	// pizz.PrintPizza()
	// pizz.PrintSlices()

	slicer.FindSlice(&pizz)

	// pizz.PrintPizzaCells()
	pizz.PrintSlices()
	pizz.PrintSlicesPlain()
	pizz.PrintScore()
}
