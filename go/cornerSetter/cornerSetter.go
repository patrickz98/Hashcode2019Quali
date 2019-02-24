package cornerSetter

import (
	"../pizza"
	"../slicer"
	// "os"
)

type CornerSetter struct {
	Slicer *slicer.Slicer
	Pizza  *pizza.Pizza
	Params CornerParams

	NeuralNet *NeuralNet
}

func (cornerSetter *CornerSetter) SetSlices() bool {

	addedSomething := false

	for xy, slices := range cornerSetter.Slicer.TopLeftSliceCache {

		if cornerSetter.Pizza.Cells[xy].Slice != nil {
			continue
		}

		if len(slices) == 0 {
			continue
		}

		addedSlice := slices[0]
		highscore := float32(0)

		for _, sl := range slices {

			contenderScore := cornerSetter.sliceValueScorer(xy, sl, Min(xy.Row, xy.Row-cornerSetter.Pizza.Row.End), Min(xy.Column, xy.Column-cornerSetter.Pizza.Column.End))

			if highscore < contenderScore {
				highscore = contenderScore
				addedSlice = sl
			}
		}

		if highscore >= cornerSetter.Params.CornerMinScore {
			if cornerSetter.Pizza.SafeAddSlice(addedSlice) {
				addedSomething = true
			}
		}
	}

	return addedSomething
}
