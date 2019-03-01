package main

import (
	"../simple"
	"fmt"
	"io/ioutil"
	"strings"
)

type SlideShow struct {
	Photos         []*Photo
	Slides         []*Slide
	params         SlideParams
	interestFactor int
}

func (this SlideShow) InterestFactor() int {

	score := 0

	for inx := 0; inx < len(this.Slides) - 1; inx++ {

		S1 := this.Slides[ inx ].Tags()
		S2 := this.Slides[ inx + 1 ].Tags()

		common  := S1.Intersection(S2)
		S1NotS2 := S1.Difference(S2)
		S2NotS1 := S2.Difference(S1)

		fmt.Println(common)
		fmt.Println(S1NotS2)
		fmt.Println(S2NotS1)

		transScore := simple.Min(
			common.Len(),
			S1NotS2.Len(),
			S2NotS1.Len())

		score += transScore
	}

	this.interestFactor = score

	return score
}

func (this SlideShow) Submission() {

	count := len(this.Slides)
	submission := make([][]int, count)
	submissionStr := fmt.Sprintf("%d\n", count)

	for inx, slide := range this.Slides {

		ids := slide.PhotoIDs()
		submission[ inx ] = ids

		str := fmt.Sprintln(ids)
		str = strings.ReplaceAll(str, "[", "")
		str = strings.ReplaceAll(str, "]", "")

		submissionStr += str
	}

	bytes := []byte(submissionStr)
	err := ioutil.WriteFile(this.params.SubmissionPath(), bytes, 0644)
	simple.CheckErr(err)

	bytes = []byte(fmt.Sprintf("Score: %d\n", this.interestFactor))
	err = ioutil.WriteFile(this.params.ScorePath(), bytes, 0644)
	simple.CheckErr(err)
}
