package main

import "fmt"

func main()  {

	fmt.Println("Start")

	inputDir := "../../input/"
	submissionDir := "../../submissions/"

	params := SlideParams{
		InputPath: inputDir + "a_example.txt",
		SubmissionDir: submissionDir,
	}

	//params := SlideParams{
	//	InputPath: inputDir + "b_lovely_landscapes.txt",
	//	SubmissionDir: submissionDir,
	//}

	//params := SlideParams{
	//	InputPath: inputDir + "c_memorable_moments.txt",
	//	SubmissionDir: submissionDir,
	//}

	//params := SlideParams{
	//	InputPath: inputDir + "d_pet_pictures.txt",
	//	SubmissionDir: submissionDir,
	//}

	//params := SlideParams{
	//	InputPath: inputDir + "e_shiny_selfies.txt",
	//	SubmissionDir: submissionDir,
	//}

	slideshow := Init(params)

	S1 := &Slide{
		Photos: []*Photo{slideshow.Photos[ 0 ]},
	}

	S2 := &Slide{
		Photos: []*Photo{slideshow.Photos[ 3 ]},
	}

	S3 := &Slide{
		Photos: []*Photo{slideshow.Photos[ 1 ], slideshow.Photos[ 2 ]},
	}

	slideshow.Slides = []*Slide{S1, S2, S3}
	slideshow.InterestFactor()
	slideshow.Submission()
}
