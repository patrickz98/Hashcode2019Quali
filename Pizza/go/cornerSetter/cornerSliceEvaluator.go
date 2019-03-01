package cornerSetter

import (
	"../pizza"
	// "os"
)

type NeuralNet struct {
	NodeCount []int

	InputLayer []float32
	Layer      [][]float32
	Output     float32

	UseJFMethod bool
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
	//neuralNet.NodeCount = []int{17, 10, 6, 5, 3, 1}
	neuralNet.NodeCount = []int{7, 5, 5, 3, 1}
	neuralNet.InputLayer = make([]float32, 7)

	neuralNet.Layer = make([][]float32, 0)
	for i := 0; i < len(neuralNet.NodeCount); i++ {
		neuralNet.Layer = append(neuralNet.Layer, make([]float32, neuralNet.NodeCount[i]))
	}
}

func (neuralNet *NeuralNet) ComputeOutput(params CornerParams) {

	if neuralNet.UseJFMethod {
		for true {
			change := false
			for i := 0; i < len(neuralNet.InputLayer)-1; i++ {
				if params.InputOrder[i] > params.InputOrder[i+1] {
					change = true
					params.InputOrder[i+1], params.InputOrder[i] = params.InputOrder[i], params.InputOrder[i+1]
					params.InputTarget[i+1], params.InputTarget[i] = params.InputTarget[i], params.InputTarget[i+1]
					neuralNet.InputLayer[i+1], neuralNet.InputLayer[i] = neuralNet.InputLayer[i], neuralNet.InputLayer[i+1]
				}
			}

			if !change {
				break
			}
		}

		//fmt.Println(params.InputOrder, params.InputTarget, neuralNet.InputLayer)

		neuralNet.Layer[0] = make([]float32, neuralNet.NodeCount[0])
		for i := 0; i < len(neuralNet.InputLayer); i++ {
			neuralNet.Layer[0][params.InputTarget[i]] = neuralNet.InputLayer[i]
		}
	} else {
		for i := 0; i < len(neuralNet.InputLayer); i++ {
			neuralNet.Layer[0][i] = neuralNet.InputLayer[i]
		}
	}
	pos := 0
	for i, _ := range neuralNet.NodeCount {
		if i == len(neuralNet.NodeCount)-1 {
			continue
		}

		for j := 0; j < neuralNet.NodeCount[i+1]; j++ {
			neuralNet.Layer[i+1][j] = 0
			divisor := float32(0)

			for k := 0; k < neuralNet.NodeCount[i]; k++ {
				neuralNet.Layer[i+1][j] += neuralNet.Layer[i][k] * params.NeuralNet[pos]
				divisor += params.NeuralNet[pos]
				pos += 1
			}

			if divisor != 0 {
				neuralNet.Layer[i+1][j] = MinMax(-1, 1, neuralNet.Layer[i+1][j]/divisor)
			}
		}
	}

	neuralNet.Output = neuralNet.Layer[len(neuralNet.NodeCount)-1][0]
}

func (cornerSetter *CornerSetter) sliceValueScorer(xy pizza.Coordinate, slice *pizza.Slice, iteration int, rowDepth int, columnDepth int) float32 {

	const PreNeighbours = 7

	cornerSetter.NeuralNet.InputLayer[0] = float32(slice.Size()-(cornerSetter.Pizza.Ingredients*2)) / float32(cornerSetter.Pizza.MaxCells-(cornerSetter.Pizza.Ingredients*2))

	cornerSetter.NeuralNet.InputLayer[1] = float32(rowDepth) / float32(cornerSetter.Pizza.Row.End)
	cornerSetter.NeuralNet.InputLayer[2] = float32(columnDepth) / float32(cornerSetter.Pizza.Column.End)

	cornerSetter.NeuralNet.InputLayer[3] = float32(len(cornerSetter.Slicer.SliceCache[xy])) / float32(MaxSlicesInAPlace)
	_, cornerSetter.NeuralNet.InputLayer[4] = cornerSetter.Pizza.Score()
	cornerSetter.NeuralNet.InputLayer[5] = float32(iteration) / float32(5)

	cornerSetter.NeuralNet.InputLayer[6] = MinMax(0, float32(1), float32(cornerSetter.SetLastSlice)/float32(200))

	//fmt.Println(cornerSetter.NeuralNet.InputLayer)

	/*cornerSetter.NeuralNet.InputLayer[PreNeighbours] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(1, -1), true, false)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+1] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(1, -1), false, false)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+2] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(1, slice.Column.Length()-1), false, false)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+3] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length()-1, -1), true, false)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+4] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(0, slice.Column.Length()), true, true)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+5] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length(), 0), true, true)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+6] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(0, slice.Column.Length()), false, true)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+7] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length(), 0), false, true)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+8] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length(), slice.Column.Length()-1), false, false)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+9] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length()-1, slice.Column.Length()), true, false)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+10] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length(), slice.Column.Length()), false, true)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+11] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length(), slice.Column.Length()), true, true)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+12] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(1, slice.Column.Length()-1), false, false)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+13] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length()-1, -1), true, false)

	cornerSetter.NeuralNet.InputLayer[PreNeighbours+14] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length(), slice.Column.Length()-1), false, false)
	cornerSetter.NeuralNet.InputLayer[PreNeighbours+15] = cornerSetter.GetFittingSliceCountFromCache(xy.AddTo(slice.Row.Length()-1, slice.Column.Length()), true, false)*/

	cornerSetter.NeuralNet.ComputeOutput(cornerSetter.Params)

	//fmt.Println(cornerSetter.NeuralNet.Output)

	return cornerSetter.NeuralNet.Output
}

func (cornerSetter *CornerSetter) GetFittingSliceCountFromCache(xy pizza.Coordinate, row bool, start bool) float32 {
	if xy.Row < 0 || xy.Column < 0 ||
		xy.Row >= cornerSetter.Slicer.Pizza.Row.End || xy.Column >= cornerSetter.Pizza.Column.End {
		return float32(1)
	}

	if row && start {
		return cornerSetter.Cache[0][xy.GetTransPos(*cornerSetter.Pizza)]
	} else if row && !start {
		return cornerSetter.Cache[1][xy.GetTransPos(*cornerSetter.Pizza)]
	} else if !row && start {
		return cornerSetter.Cache[2][xy.GetTransPos(*cornerSetter.Pizza)]
	} else {
		return cornerSetter.Cache[3][xy.GetTransPos(*cornerSetter.Pizza)]
	}
}
