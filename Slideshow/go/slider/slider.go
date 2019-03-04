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

func (this *Slider) findVertical() {

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
}

func (this *Slider) findVCouples() {

	//couples := make([][]*show.Photo, 0)
	couples := set.New()

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
	fmt.Println("couples:", couples.Len())
}

func (this *Slider) merge(slides1 []*show.Slide, slides2 []*show.Slide) {


}

func (this *Slider) find() {


}
