package show

import (
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

func (this Slide) PhotoIDs() []int {

	result := make([]int, 0)

	for _, photo := range this.Photos {

		result = append(result, photo.ID)
	}

	return result
}
