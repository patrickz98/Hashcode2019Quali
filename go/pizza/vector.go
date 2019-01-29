package pizza

import (
	"strconv"
)

type Vector struct {
	Start int
	End   int
}

func (vec Vector) Length() int {
	return vec.End - vec.Start + 1
}

func (vec Vector) CutPossible() bool {
	length := vec.End - vec.Start

	return length > 0
}

func (vec Vector) Cut() (Vector, Vector) {

	length := vec.End - vec.Start
	mid := vec.Start + length/2
	v1 := Vector{Start: vec.Start, End: mid}
	v2 := Vector{Start: mid + 1, End: vec.End}

	return v1, v2
}

func (vec Vector) Stringify() string {

	return "(" + strconv.FormatInt(int64(vec.Start), 10) + ", " +
		strconv.FormatInt(int64(vec.End), 10) + ")"
}

func (vec Vector) Range() []int {

	// fmt.Printf("vec.Length = %d\n", vec.Length())
	numbers := make([]int, vec.Length())

	inx := 0
	for num := vec.Start; num <= vec.End; num++ {
		numbers[ inx ] = num
		inx++
	}

	return numbers
}

func (vec Vector) Size(vec1 Vector) int {

	return vec.Length() * vec1.Length()
}