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
	neuralNet.NodeCount = []int{5, 5, 5, 3, 1}

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

	cornerSetter.NeuralNet.Layer[0][0] = float32(slice.Size()-(cornerSetter.Pizza.Ingredients*2)) / float32(cornerSetter.Pizza.MaxCells-(cornerSetter.Pizza.Ingredients*2))

	cornerSetter.NeuralNet.Layer[0][1] = float32(rowDepth) / float32(cornerSetter.Pizza.Row.End)
	cornerSetter.NeuralNet.Layer[0][2] = float32(columnDepth) / float32(cornerSetter.Pizza.Column.End)

	cornerSetter.NeuralNet.Layer[0][3] = float32(len(cornerSetter.Slicer.SliceCache[xy])) / float32(MaxSlicesInAPlace)
	_, cornerSetter.NeuralNet.Layer[0][4] = cornerSetter.Pizza.Score()

	cornerSetter.NeuralNet.ComputeOutput(&cornerSetter.Params.NeuralNet)

	return cornerSetter.NeuralNet.Output
}
