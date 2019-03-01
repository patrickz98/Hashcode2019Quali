package show

import (
	"fmt"
	"github.com/golang-collections/collections/set"
)

type Photo struct {
	ID          int
	Orientation string
	Tags        *set.Set
	TagsLen     int
}

func (this Photo) Horizontal() bool {

	return this.Orientation == "H"
}

func (this Photo) Vertical() bool {

	return this.Orientation == "V"
}

func (this Photo) Print() {

	tags := make([]string, this.Tags.Len())

	inx := 0
	this.Tags.Do(func(val interface{}) {

		tags[ inx ] = val.(string)
		inx++
	})

	fmt.Printf("ID: %v\n", this.ID)
	fmt.Printf("Orientation: %v\n", this.Orientation)
	fmt.Printf("Tags: %v\n", tags)
	fmt.Printf("TagsLen: %v\n", this.TagsLen)
}
