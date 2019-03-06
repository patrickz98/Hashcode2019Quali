package main

import (
	"../show"
	"../slider"
	"fmt"
)

func main()  {

	//test := make([]int, 4)
	//test[0] = 1
	//test[1] = 2
	//test[2] = 3
	//test[3] = 4
	//
	//part1 := append([]int{}, test[:1]...)
	//part2 := append([]int{}, test[1:]...)
	//
	//fmt.Println("part1", part1)
	//fmt.Println("part2", part2)
	//
	//tmp := append(part1, 0)
	//tmp = append(tmp, part2...)
	//fmt.Println("tmp", tmp)
	//
	//os.Exit(0)

	fmt.Println("Start")

	inputDir := "../../input/"
	submissionDir := "../../submissions/"

	params := show.SlideParams{
		// Verticals: 2 Horizontals: 2 Couples: 1
		//InputPath: inputDir + "a_example.txt",

		// Verticals: 0 Horizontals: 80000 Couples: 0
		InputPath: inputDir + "b_lovely_landscapes.txt",

		// Verticals: 500 Horizontals: 500 Couples: 124750
		//InputPath: inputDir + "c_memorable_moments.txt",

		// Verticals: 60000 Horizontals: 30000 V-Couples: 1799970000
		//InputPath: inputDir + "d_pet_pictures.txt",

		// Verticals: 80000 Horizontals: 0 Couples: 3199960000
		//InputPath: inputDir + "e_shiny_selfies.txt",
		SubmissionDir: submissionDir,
	}

	slideshow := show.Init(params)

	slider.Find(slideshow)

	//slideshow.InterestFactor()
	slideshow.Submission()
}