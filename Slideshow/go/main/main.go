package main

import "fmt"

func main()  {

	fmt.Println("Start")

	inputPath := "../../input/a_example.txt"
	//inputPath := "../../input/b_lovely_landscapes.txt"
	//inputPath := "../../input/c_memorable_moments.txt"
	//inputPath := "../../input/d_pet_pictures.txt"
	//inputPath := "../../input/e_shiny_selfies.txt"

	slideshow := Init(inputPath)

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
