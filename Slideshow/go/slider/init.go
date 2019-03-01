package slider

import (
	"../show"
	"fmt"
)

func Find(slideShow *show.SlideShow) {

	slider := Slider{Show: slideShow}
	slider.findVertical()
	slider.find()

	fmt.Println("Interest-Factor: ", slideShow.InterestFactor())
}