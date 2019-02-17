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

	// 99.08%
	inputPath := "../../input/c_medium.in"
	submissionPath := "../../submissions/c_medium.out"

<<<<<<< HEAD
	// inputPath := "../../input/a_example.in"
	//inputPath := "../../input/b_small.in"
	// inputPath := "../../input/c_medium.in"
	 inputPath := "../../input/d_big.in"
=======
	// 93.06%
	// inputPath := "../../input/d_big.in"
	// submissionPath := "../../submissions/d_big.out"
>>>>>>> 1481eba42e4ac93dd1e584fac1a757b3e2c2e742

	piz := pizza.NewPizza(inputPath)

	// piz = pizza.Pizza{
	// 	Ingredients: pizz.Ingredients,
	// 	MaxCells:    pizz.MaxCells,
	// 	Cells:       pizz.Cells,
	// 	Row: pizza.Vector{Start: 0, End: 99},
	// 	Column: pizza.Vector{Start: 0, End: 99},
	// }

	piz.PrintParams()

	slicer.SearchSlices(&piz)
	// piz.CheckErrors()

	// piz.PrintSlices(false)
	piz.CreateSubmission(submissionPath)
	piz.PrintScore()

	piz.PrintSlicesToFile(true, "xxx-marked.txt")
	piz.PrintSlicesToFile(false, "xxx.txt")

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
