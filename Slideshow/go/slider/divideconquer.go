package slider

import (
	"../show"
	"fmt"
)

func (this *Slider) merge(slides1 []*show.Slide, slides2 []*show.Slide) {

	//slides1Len := len(slides1)
	//slides2Len := len(slides2)

	mergedSlides := make([]*show.Slide, 0)
	mergedSlides = append(mergedSlides, slides1...)

	for count, slide := range slides2 {

		fmt.Println("count", count)
		slide.PrintTags()

		gain := 0
		positions := 0

		for iny := 0; iny < len(mergedSlides); iny++ {

			cur := mergedSlides[ iny ]

			scoreCur := 0
			scoreInsert := 0

			if iny > 0 {
				pre := mergedSlides[ iny - 1 ]
				scoreCur += pre.InterestFactor(*cur)
				scoreInsert += pre.InterestFactor(*slide)
			}

			if iny < len(mergedSlides) - 1 {
				post := mergedSlides[ iny + 1 ]
				scoreCur += cur.InterestFactor(*post)
			}

			scoreInsert += slide.InterestFactor(*cur)

			//if scoreCur > 0 || scoreInsert > 0 {
			//	fmt.Println("################")
			//	fmt.Println("scoreCur", scoreCur)
			//	fmt.Println("scoreInsert", scoreInsert)
			//}

			if scoreCur > scoreInsert {
				continue
			}

			if gain >= scoreInsert {
				continue
			}

			gain = scoreInsert
			positions = iny

			fmt.Println("gain", gain)
		}

		part1 := mergedSlides[:positions]
		part2 := mergedSlides[positions:]

		merge := append(part1, slide)
		merge = append(merge, part2...)

		mergedSlides = merge
	}

	fmt.Println("++++++++++++++++++ mergedSlides =", this.Show.InterestFactorFor(mergedSlides))
}

func (this *Slider) findRec() {

	photos := this.Show.Photos

	slides := make([]*show.Slide, 0)

	//for _, photo := range photos {
	for _, photo := range photos[:2000] {

		slide := show.NewSlide(photo)
		slides = append(slides, slide)
	}

	fmt.Println("++++++++++++++++++ slides=", this.Show.InterestFactorFor(slides))

	mid := len(slides) / 2
	slides1 := slides[ : mid]
	slides2 := slides[ mid : ]

	//fmt.Println("slides", slides)
	//fmt.Println("slides1", slides1)
	//fmt.Println("slides2", slides2)

	this.merge(slides1, slides2)
}
