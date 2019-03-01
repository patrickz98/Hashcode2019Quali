package slider

import (
	"../show"
	"fmt"
)


type Slider struct {
	Show *show.SlideShow
}

func (this *Slider) findVertical() {

	vertical := make([]*show.Photo, 0)

	length := len(this.Show.Photos)

	for inx := 0; inx < length; inx++ {

		photo := this.Show.Photos[ inx ]

		if photo.Horizontal() {
			continue
		}

		vertical = append(vertical, photo)
	}

	//fmt.Println(couples)
	fmt.Println("verticals:", len(vertical))
}

func (this *Slider) findVCouples() {

	couples := make([][]*show.Photo, 0)

	length := len(this.Show.Photos)

	for inx := 0; inx < length; inx++ {
		fmt.Println("inx", inx)

		for iny := inx + 1; iny < length; iny++ {

			//fmt.Println("inx", inx, "iny", iny)

			photo1 := this.Show.Photos[ inx ]
			photo2 := this.Show.Photos[ iny ]

			if photo1.Horizontal() || photo2.Horizontal() {
				continue
			}

			couples = append(couples, []*show.Photo{photo1, photo2})
		}
	}

	//fmt.Println(couples)
	fmt.Println("couples:", len(couples))
}

func (this *Slider) find() {

	//photos := set.New()
	//
	//for _, photo := range this.Show.Photos {
	//	photos.Insert(*photo)
	//}

	//photos1.Do(func(val interface{}) {
	//
	//	photo := val.(show.Photo)
	//	fmt.Println("ID:", photo.ID)
	//})
	//
	//photos2 := photos1.Union(set.New())

	//for _, photo := range this.Show.Photos {
	//
	//	//fmt.Println("Has:", photos.Has(*photo))
	//}

	//S1 := &show.Slide{
	//	Photos: []*show.Photo{this.Show.Photos[ 0 ]},
	//}
	//
	//S2 := &show.Slide{
	//	Photos: []*show.Photo{this.Show.Photos[ 3 ]},
	//}
	//
	//S3 := &show.Slide{
	//	Photos: []*show.Photo{this.Show.Photos[ 1 ], this.Show.Photos[ 2 ]},
	//}

	//this.Show.Slides = []*show.Slide{S1, S2, S3}
}
