package slider

import (
	"../show"
	"fmt"
	"github.com/golang-collections/collections/set"
)

type Slider struct {
	Show *show.SlideShow
}

type coupleSync struct {
	couples           [][]*show.Slide
	leftovers         *set.Set
	interestingFactor int
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

func (this *Slider) findBestCouples(logtag string, slides *set.Set, done chan coupleSync) {

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
		fmt.Println(logtag, "totalGain", totalGain)

		couple := []*show.Slide{slide1, bestSlide}
		couples = append(couples, couple)

		slides.Remove(bestSlide)
	})

	result := coupleSync{
		couples: couples,
		leftovers: leftOvers,
		interestingFactor: totalGain,
	}

	done <- result
}

func (this *Slider) findBestPositionIn(slides []*show.Slide, place *show.Slide) (factor int, index int) {

	factor = 0
	index = -1

	for inx := 0; inx <= len(slides); inx++ {

		lost := 0
		gain := 0

		if inx > 0 && inx < len(slides) {
			pre := slides[ inx - 1 ]
			cur := slides[ inx ]

			lost += pre.InterestFactor(*cur)
		}

		// pre place
		if inx > 0 {
			slide := slides[ inx - 1 ]
			gain += slide.InterestFactor(*place)
		}

		// post place
		if inx < len(slides) {
			slide := slides[ inx ]
			gain += place.InterestFactor(*slide)
		}

		gain -= lost

		if gain > factor {
			factor = gain
			index = inx
		}
	}

	return factor, index
}

func (this *Slider) findBestInCouples(couples [][]*show.Slide, slide *show.Slide) (int, int, int) {

	bestfactor := 0
	bestindex := 0
	bestcoupleindex := 0

	for inx, couple := range couples {

		factor, index := this.findBestPositionIn(couple, slide)

		if factor <= 0 {
			continue
		}

		if factor <= bestfactor {
			continue
		}

		bestfactor = factor
		bestindex = index
		bestcoupleindex = inx
	}

	return bestfactor, bestindex, bestcoupleindex
}

func (this *Slider) InterestFactorFor(slides [][]*show.Slide) int {

	mergerd := make([]*show.Slide, 0)

	for _, slideShow := range slides {
		mergerd = append(mergerd, slideShow...)
	}

	return this.Show.InterestFactorFor(mergerd)
}

func (this *Slider) findMerge(results... coupleSync) {

	allCouples := make([][]*show.Slide, 0)
	leftovers := set.New()

	for _, result := range results {
		allCouples = append(allCouples, result.couples...)
		leftovers = leftovers.Union(result.leftovers)
	}

	fmt.Println("#### InterestFactorFor", this.InterestFactorFor(allCouples))

	count := 0
	leftoverslen := leftovers.Len()
	gain := 0

	leftovers.Do(func(val interface{}) {

		count++
		fmt.Printf("merge %4.1f%% gain %d\r", float32(count) / float32(leftoverslen) * 100, gain)

		slide := val.(*show.Slide)

		bestfactor, bestindex, bestcoupleindex := this.findBestInCouples(allCouples, slide)

		if bestfactor <= 0 {
			return
		}

		gain += bestfactor
		//fmt.Println("gain", bestfactor)
		//fmt.Println("bestindex", bestindex)
		//fmt.Println("bestcoupleindex", bestcoupleindex)

		couple := allCouples[ bestcoupleindex ]

		part1 := append([]*show.Slide{}, couple[:bestindex]...)
		part2 := append([]*show.Slide{}, couple[bestindex:]...)

		tmp := append(part1, slide)
		tmp = append(tmp, part2...)

		allCouples[ bestcoupleindex ] = tmp
	})

	fmt.Println()
	fmt.Println("#### InterestFactorFor", this.InterestFactorFor(allCouples))
}

func (this *Slider) find() {

	splitfactor := 4

	slides := make([]*set.Set, splitfactor)

	for inx := range slides {
		slides[ inx ] = set.New()
	}

	for inx, photo := range this.Show.Photos[:4000] {

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

	done := make(chan coupleSync, splitfactor)

	for inx, slide := range slides {
		logtag := fmt.Sprintf("#%d", inx)
		go this.findBestCouples(logtag, slide, done)
	}

	results := make([]coupleSync, splitfactor)

	total := 0
	totalLeft := 0

	for inx := 0; inx < splitfactor; inx++ {

		result := <- done
		results[ inx ] = result

		factor := result.interestingFactor
		total += factor
		totalLeft += result.leftovers.Len()

		fmt.Println("Result", inx, "factor", factor, "leftovers", result.leftovers.Len())
	}

	fmt.Println("total", total, "leftovers", totalLeft)

	this.findMerge(results...)
}