package show

import (
	"../simple"
	"fmt"
	"strings"
)

type SlideShow struct {
	Photos         []*Photo
	Slides         []*Slide
	params         SlideParams
}

func (this *SlideShow) InterestFactorFor(slides []*Slide) int {

	score := 0

	for inx := 0; inx < len(slides) - 1; inx++ {

		S1 := slides[ inx ]
		S2 := slides[ inx + 1 ]

		score += S1.InterestFactor(*S2)
	}

	return score
}

func (this *SlideShow) InterestFactor() int {

	return this.InterestFactorFor(this.Slides)
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

	simple.Write(this.params.SubmissionPath(), submissionStr)

	score := fmt.Sprintf("Score: %d\n", this.InterestFactor())
	simple.Write(this.params.ScorePath(), score)

	fmt.Println("Submission", this.params.SubmissionPath())
}
