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
		//	InputPath: inputDir + "a_example.txt",
		//	InputPath: inputDir + "b_lovely_landscapes.txt",
		InputPath: inputDir + "c_memorable_moments.txt",
		//	InputPath: inputDir + "d_pet_pictures.txt",
		//	InputPath: inputDir + "e_shiny_selfies.txt",
		SubmissionDir: submissionDir,
	}

	slideshow := show.Init(params)

	slider.Find(slideshow)

	//slideshow.InterestFactor()
	slideshow.Submission()
}