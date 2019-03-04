package main

import (
	"../show"
	"../slider"
	"fmt"
)

func main()  {

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