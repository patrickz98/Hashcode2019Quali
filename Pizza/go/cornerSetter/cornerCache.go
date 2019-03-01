package cornerSetter

import (
	"../pizza"
	// "os"
)

func (cornerTrainer *CornerTrainer) buildSlicesCache() {
	cache := make([][]float32, 4)

	for i := 0; i < 4; i++ {
		cache[i] = make([]float32, cornerTrainer.Slicer.Pizza.Size())
	}

	cornerSetter := cornerTrainer.CornerSetter

	for count, xy := range cornerTrainer.Slicer.Pizza.Traversal() {
		cache[0][count] = cornerSetter.GetFittingSliceCount(xy, true, true)
		cache[1][count] = cornerSetter.GetFittingSliceCount(xy, true, false)

		cache[2][count] = cornerSetter.GetFittingSliceCount(xy, false, true)
		cache[3][count] = cornerSetter.GetFittingSliceCount(xy, false, false)
	}

	cornerTrainer.CornerSetter.Cache = cache
}

func (cornerSetter *CornerSetter) GetFittingSliceCount(xy pizza.Coordinate, row bool, start bool) float32 {

	if xy.Row < 0 || xy.Column < 0 ||
		xy.Row >= cornerSetter.Slicer.Pizza.Row.End || xy.Column >= cornerSetter.Pizza.Column.End ||
		len(cornerSetter.Slicer.SliceCache[xy]) == 0 {
		return float32(1)
	}

	count := 0
	for _, slice := range cornerSetter.Slicer.SliceCache[xy] {
		if row && start && slice.Row.Start == xy.Row {
			count++
		} else if row && !start && slice.Row.End == xy.Row {
			count++
		} else if !row && start && slice.Column.Start == xy.Column {
			count++
		} else if !row && !start && slice.Column.End == xy.Column {
			count++
		}
	}

	return float32(count) / float32(len(cornerSetter.Slicer.TopLeftSliceCache[xy]))
}
