package cornerSetter

import (
	"../pizza"
	// "os"
)

func (cornerSetter *CornerSetter) sliceValueScorer(slice *pizza.Slice, depth int) float32 {

	return float32(slice.Size()) / (float32(depth) * 0.025)
}
