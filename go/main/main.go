package main

import (
	"../pizza"
	"../slicer"
	"fmt"
)

func printParts(parts []*pizza.Pizza) {

	for _, parts := range parts {
		if parts != nil {
			fmt.Println("--------------")
			parts.PrintPizza()
		}
	}
}

func main() {

	fmt.Println("Start")

	// vec := pizza.Vector{Start: 2, End: 4}
	//
	// for _, inx := range vec.Range() {
	// 	fmt.Printf("inx=%d\n", inx)
	// }

	// inputPath := "../../input/a_example.in"
	// inputPath := "../../input/b_small.in"
	// inputPath := "../../input/c_medium.in"
	inputPath := "../../input/d_big.in"

	pizz := pizza.NewPizza(inputPath)
	pizz.PrintParams()
	// pizz.PrintPizza()
	fmt.Println("-------")

	slicer.SearchSlices(&pizz)

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
