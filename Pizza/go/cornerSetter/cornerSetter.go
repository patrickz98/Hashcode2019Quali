package cornerSetter

import (
	"../pizza"
	"../slicer"
	// "os"
)

type CornerSetter struct {
	Slicer       *slicer.Slicer
	Pizza        *pizza.Pizza
	Params       CornerParams
	SetLastSlice int
	Cache        [][]float32

	NeuralNet *NeuralNet
}

func (cornerSetter *CornerSetter) SetSlices(currentIteration int, settingThreshold float32) bool {

	addedSomething := false

	for xy, slices := range cornerSetter.Slicer.TopLeftSliceCache {

		if len(slices) == 0 {
			continue
		}

		addedSlice := slices[0]
		highscore := float32(0)

		for _, sl := range slices {

			contenderScore := cornerSetter.sliceValueScorer(xy, sl, currentIteration, Min(xy.Row, cornerSetter.Pizza.Row.End-xy.Row), Min(xy.Column, cornerSetter.Pizza.Column.End-xy.Column))

			if highscore < contenderScore {
				highscore = contenderScore
				addedSlice = sl
			}
		}

		if highscore >= settingThreshold {
			if cornerSetter.Pizza.SafeAddSlice(addedSlice) {
				addedSomething = true
				cornerSetter.SetLastSlice = xy.GetTransPos(*cornerSetter.Pizza)
			}
		}
	}

	return addedSomething
}
