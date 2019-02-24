package cornerSetter

import (
	"../pizza"
	// "os"
)

type NeuralNet struct {
	NodeCount []int

	Layer  [][]float32
	Output float32
}

func (neuralNet *NeuralNet) TotalConnectionsCount() int {
	total := 0

	for i := 1; i < len(neuralNet.NodeCount); i++ {
		total += neuralNet.NodeCount[i-1] * neuralNet.NodeCount[i]
	}

	return total
}

func (neuralNet *NeuralNet) Init() {
	//neuralNet.NodeCount = []int{5, 5, 5, 3, 1}
	neuralNet.NodeCount = []int{17, 10, 6, 5, 3, 1}

	neuralNet.Layer = make([][]float32, 0)
	for i := 0; i < len(neuralNet.NodeCount); i++ {
		neuralNet.Layer = append(neuralNet.Layer, make([]float32, neuralNet.NodeCount[i]))
	}
}

func (neuralNet *NeuralNet) ComputeOutput(data *[]float32) {
	pos := 0

	for i, _ := range neuralNet.NodeCount {
		if i == len(neuralNet.NodeCount)-1 {
			continue
		}

		for j := 0; j < neuralNet.NodeCount[i+1]; j++ {
			neuralNet.Layer[i+1][j] = 0
			divisor := float32(0)

			for k := 0; k < neuralNet.NodeCount[i]; k++ {
				neuralNet.Layer[i+1][j] += neuralNet.Layer[i][k] * (*data)[pos]
				divisor += (*data)[pos]
				pos += 1
			}

			neuralNet.Layer[i+1][j] = neuralNet.Layer[i+1][j] / divisor
		}
	}

	neuralNet.Output = neuralNet.Layer[len(neuralNet.NodeCount)-1][0]
}

func (cornerSetter *CornerSetter) sliceValueScorer(xy pizza.Coordinate, slice *pizza.Slice, rowDepth int, columnDepth int) float32 {

	const PreNeighbours = 5

	cornerSetter.NeuralNet.Layer[0][0] = float32(slice.Size()-(cornerSetter.Pizza.Ingredients*2)) / float32(cornerSetter.Pizza.MaxCells-(cornerSetter.Pizza.Ingredients*2))

	cornerSetter.NeuralNet.Layer[0][1] = float32(rowDepth) / float32(cornerSetter.Pizza.Row.End)
	cornerSetter.NeuralNet.Layer[0][2] = float32(columnDepth) / float32(cornerSetter.Pizza.Column.End)

	cornerSetter.NeuralNet.Layer[0][3] = float32(len(cornerSetter.Slicer.SliceCache[xy])) / float32(MaxSlicesInAPlace)
	_, cornerSetter.NeuralNet.Layer[0][4] = cornerSetter.Pizza.Score()

	for i := 0; i < 12; i++ {
		cornerSetter.NeuralNet.Layer[0][PreNeighbours+i] = 0
	}

	cornerSetter.NeuralNet.Layer[0][PreNeighbours] = cornerSetter.GetFittingSliceCount(xy.AddTo(1, -1), true, false)
	cornerSetter.NeuralNet.Layer[0][PreNeighbours+1] = cornerSetter.GetFittingSliceCount(xy.AddTo(1, -1), false, false)

	cornerSetter.NeuralNet.Layer[0][PreNeighbours+2] = cornerSetter.GetFittingSliceCount(xy.AddTo(1, slice.Column.Length()-1), false, false)
	cornerSetter.NeuralNet.Layer[0][PreNeighbours+3] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length()-1, -1), true, false)

	cornerSetter.NeuralNet.Layer[0][PreNeighbours+4] = cornerSetter.GetFittingSliceCount(xy.AddTo(0, slice.Column.Length()), true, true)
	cornerSetter.NeuralNet.Layer[0][PreNeighbours+5] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length(), 0), true, true)

	cornerSetter.NeuralNet.Layer[0][PreNeighbours+6] = cornerSetter.GetFittingSliceCount(xy.AddTo(0, slice.Column.Length()), false, true)
	cornerSetter.NeuralNet.Layer[0][PreNeighbours+7] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length(), 0), false, true)

	cornerSetter.NeuralNet.Layer[0][PreNeighbours+8] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length(), slice.Column.Length()-1), false, false)
	cornerSetter.NeuralNet.Layer[0][PreNeighbours+9] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length()-1, slice.Column.Length()), true, false)

	cornerSetter.NeuralNet.Layer[0][PreNeighbours+10] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length(), slice.Column.Length()), false, true)
	cornerSetter.NeuralNet.Layer[0][PreNeighbours+11] = cornerSetter.GetFittingSliceCount(xy.AddTo(slice.Row.Length(), slice.Column.Length()), true, true)

	cornerSetter.NeuralNet.ComputeOutput(&cornerSetter.Params.NeuralNet)

	return cornerSetter.NeuralNet.Output
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
