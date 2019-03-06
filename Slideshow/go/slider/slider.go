package slider

import (
	"../show"
	"fmt"
	"github.com/golang-collections/collections/set"
)

type Slider struct {
	Show *show.SlideShow
}

func (this *Slider) statistics() {

	verticals := 0
	horizontals := 0

	length := len(this.Show.Photos)

	for inx := 0; inx < length; inx++ {

		photo := this.Show.Photos[ inx ]

		if photo.Horizontal() {
			horizontals++
		} else {
			verticals++
		}
	}

	fmt.Println("verticals:", verticals)
	fmt.Println("horizontals:", horizontals)
}

func (this *Slider) findVertical() []*show.Photo {

	verticals := make([]*show.Photo, 0)

	length := len(this.Show.Photos)

	for inx := 0; inx < length; inx++ {

		photo := this.Show.Photos[ inx ]

		if photo.Horizontal() {
			continue
		}

		verticals = append(verticals, photo)
	}

	//fmt.Println(couples)
	fmt.Println("verticals:", len(verticals))

	return verticals
}

func (this *Slider) findVerticalCouples() {

	//couples := make([][]*show.Photo, 0)
	//couples := set.New()

	length := len(this.Show.Photos)

	count := 0

	for inx := 0; inx < length; inx++ {
		fmt.Printf("inx=%d count=%d\n", inx, count)

		for iny := inx + 1; iny < length; iny++ {

			//fmt.Println("inx", inx, "iny", iny)

			photo1 := this.Show.Photos[ inx ]
			photo2 := this.Show.Photos[ iny ]

			if photo1.Horizontal() || photo2.Horizontal() {
				continue
			}

			//slide := show.NewSlide(photo1, photo2)
			//slide := []*show.Photo{photo1, photo2}
			//couples.Insert(slide)

			count++
			//couples = append(couples, []*show.Photo{photo1, photo2})
		}
	}

	//fmt.Println(couples)
	fmt.Println("count:", count)
	//fmt.Println("couples:", len(couples))
	//fmt.Println("couples:", couples.Len())
}

func (this *Slider) findBestCouples(logtag string, slides *set.Set, done chan int) {

	couples := make([][]*show.Slide, 0)

	leftOvers := set.New()

	totalGain := 0

	slides.Do(func(val1 interface{}) {

		slide1 := val1.(*show.Slide)
		slides.Remove(slide1)

		//fmt.Println(logtag, "slides =", slides.Len())

		bestGain := 0
		var bestSlide *show.Slide

		slides.Do(func(val2 interface{}) {

			if bestSlide != nil {
				return
			}

			slide2 := val2.(*show.Slide)

			gain := slide1.InterestFactor(*slide2)

			if gain <= 0 {
				return
			}

			if bestGain < gain {
				bestGain = gain
				bestSlide = slide2
			}
		})

		if bestSlide == nil {

			leftOvers.Insert(slide1)
			return
		}

		totalGain += bestGain
		fmt.Println(logtag, "totalGain =", totalGain)

		couple := []*show.Slide{slide1, bestSlide}
		couples = append(couples, couple)

		slides.Remove(bestSlide)
	})

	//fmt.Println(logtag, "couples =", len(couples))

	commitSlices := make([]*show.Slide, 0)

	for _, couple := range couples {
		commitSlices = append(commitSlices, couple...)
	}

	//fmt.Println(logtag, "interest factor:", this.Show.InterestFactorFor(commitSlices))
	done <- this.Show.InterestFactorFor(commitSlices)
}

func (this *Slider) find() {

	splitfactor := 8

	slides := make([]*set.Set, splitfactor)

	for inx := range slides {
		slides[ inx ] = set.New()
	}

	for inx, photo := range this.Show.Photos {

		if photo.Vertical() {
			continue
		}

		slide := show.NewSlide(photo)

		index := inx % splitfactor
		slides[ index ].Insert(slide)
	}

	for inx, slide := range slides {
		fmt.Println(inx, "slide:", slide.Len())
	}

	done := make(chan int, splitfactor)

	for inx, slide := range slides {
		logtag := fmt.Sprintf("#%d", inx)
		go this.findBestCouples(logtag, slide, done)
	}

	results := make([]int, splitfactor)

	total := 0

	for inx := 0; inx < splitfactor; inx++ {
		factor := <- done
		results[ inx ] = factor

		total += factor
		fmt.Println("Result", inx, ":", factor, "total =", total)
	}

	fmt.Println("total:", total)
}