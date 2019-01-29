package pizza

import "fmt"

type Slice struct {
	Row    Vector
	Column Vector
}

func (slice Slice) Size() int {

	return slice.Row.Size(slice.Column)
}

func (slice Slice) PrintVector() {
	fmt.Printf("row=%s column=%s\n", slice.Row.Stringify(), slice.Column.Stringify())
}
