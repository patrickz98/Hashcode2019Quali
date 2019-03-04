package show

import (
	"fmt"
	"../simple"
	"github.com/golang-collections/collections/set"
)

type Slide struct {
	Photos []*Photo
	tags   *set.Set
}

func (this *Slide) Tags() *set.Set {

	if this.tags != nil {
		return this.tags
	}

	tags := set.New()

	for _, photo := range this.Photos {
		tags = tags.Union(photo.Tags)
	}

	this.tags = tags

	return this.tags
}

func (this *Slide) PrintTags() {

	this.Tags().Do(func(val interface{}) {

		fmt.Print(val, " ")
	})

	fmt.Println()
}

func (this Slide) PhotoIDs() []int {

	result := make([]int, 0)

	for _, photo := range this.Photos {

		result = append(result, photo.ID)
	}

	return result
}

func (this Slide) InterestFactor(slide Slide) int {

	S1 := this.Tags()
	S2 := slide.Tags()

	common  := S1.Intersection(S2)
	S1NotS2 := S1.Difference(S2)
	S2NotS1 := S2.Difference(S1)

	//fmt.Println(common)
	//fmt.Println(S1NotS2)
	//fmt.Println(S2NotS1)

	transScore := simple.Min(
		common.Len(),
		S1NotS2.Len(),
		S2NotS1.Len())

	return transScore
}

func NewSlide(photos ...*Photo) *Slide {

	return &Slide{Photos: photos}
}