package cornerSetter

import (
	"../slicer"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

const POPULATION = 25
const KEEP_BEST_COUNT = 5
const MUTATION_THRESHOLD = 0.5

var MaxSlicesInAPlace = 0

type CornerTrainer struct {
	Slicer       *slicer.Slicer
	ParamsLimit  CornerParamsLimit
	Generation   int
	HighScore    float32
	OldHighScore float32
	NeuralNet    NeuralNet

	CornerSetter CornerSetter
	Params       []CornerParams
	Scores       []float32
}

func (cornerTrainer *CornerTrainer) Init(path string) {
	MaxSlicesInAPlace = 0
	for _, slices := range cornerTrainer.Slicer.SliceCache {
		if len(slices) > MaxSlicesInAPlace {
			MaxSlicesInAPlace = len(slices)
		}
	}

	data, err := ioutil.ReadFile(path)

	cornerTrainer.NeuralNet = NeuralNet{}
	cornerTrainer.NeuralNet.Init()
	cornerTrainer.CornerSetter = CornerSetter{cornerTrainer.Slicer, cornerTrainer.Slicer.Pizza, CornerParams{}, &cornerTrainer.NeuralNet}
	cornerTrainer.Params = make([]CornerParams, 0)
	cornerTrainer.Scores = make([]float32, 0)

	cornerTrainer.ParamsLimit = CornerParamsLimit{
		LimitFloat32{0, 1},
		LimitFloat32NNSlice{cornerTrainer.NeuralNet.TotalConnectionsCount()},
	}

	if err == nil {
		input := strings.Split(string(data), "\n")

		firstLine := strings.Split(input[0], ",")
		gen, _ := strconv.Atoi(firstLine[0])
		cornerTrainer.Generation = gen

		for _, line := range input[1:] {
			newParams := CornerParams{}
			newParams.Init(strings.Split(line, ","))
			cornerTrainer.Params = append(cornerTrainer.Params, newParams)
		}

		fmt.Printf("Successfully read in %d lines to the CornerTrainer", len(input))
	} else {
		for i := 0; i < POPULATION; i++ {
			cornerTrainer.Params = append(cornerTrainer.Params, cornerTrainer.ParamsLimit.getRandomParams())
		}
	}
}

func (cornerTrainer *CornerTrainer) ExpandThroughCorners() {

	averageScore := float32(0)
	cornerTrainer.Scores = make([]float32, POPULATION)
	LastImprovementGen := 0
	for cornerTrainer.HighScore < 1 {
		scoreSum := float32(0)
		HighestRep := 0
		for i, params := range cornerTrainer.Params {

			addToScore := cornerTrainer.Scores[i]

			if i >= KEEP_BEST_COUNT {
				cornerTrainer.CornerSetter.Params = params

				cornerTrainer.CornerSetter.Pizza.RemoveAllSlice()

				for j := 0; j < 20; j++ {
					if j > HighestRep {
						HighestRep = j
					}

					if !cornerTrainer.CornerSetter.SetSlices() {
						break
					}
				}

				_, addToScore = cornerTrainer.CornerSetter.Pizza.Score()
			}
			//slicer.FindBiggestParts()
			//slicer.FindSingles()

			//slicer.ExpandThroughEdge()
			//cornerTrainer.CornerSetter.Slicer.FindSmallestParts()
			//cornerTrainer.CornerSetter.Slicer.ExpandThroughDestruction()
			//cornerTrainer.CornerSetter.Slicer.ExpandThroughShrink()

			scoreSum += addToScore
			cornerTrainer.Scores[i] = addToScore
			fmt.Printf("Current Generation: %d/%d, Score %.3f, AverageScore %.3f, Highscore: %.3f, Improvement: %.3f/%d, HighestRep %d,             \r", i, cornerTrainer.Generation, 100*addToScore, averageScore, 100*cornerTrainer.HighScore, 100*(cornerTrainer.HighScore-cornerTrainer.OldHighScore), LastImprovementGen, HighestRep)
		}
		averageScore = 100 * scoreSum / float32(len(cornerTrainer.Params))

		for j, _ := range cornerTrainer.Params {
			if cornerTrainer.Scores[j] > cornerTrainer.HighScore {

				cornerTrainer.OldHighScore = cornerTrainer.HighScore
				cornerTrainer.HighScore = cornerTrainer.Scores[j]
				LastImprovementGen = cornerTrainer.Generation
				//cornerTrainer.Slicer.Pizza.CreateSubmission("../../submissions_Lukas/d_big.out")
			}
		}

		cornerTrainer.Generation += 1
		cornerTrainer.AdaptParams()
	}
}

func (cornerTrainer *CornerTrainer) AdaptParams() {
	buffer := make([]CornerParams, 0)

	for true {
		change := false
		for i := 0; i < len(cornerTrainer.Scores)-1; i++ {
			if cornerTrainer.Scores[i] < cornerTrainer.Scores[i+1] {
				change = true
				cornerTrainer.Scores[i+1], cornerTrainer.Scores[i] = cornerTrainer.Scores[i], cornerTrainer.Scores[i+1]
				cornerTrainer.Params[i+1], cornerTrainer.Params[i] = cornerTrainer.Params[i], cornerTrainer.Params[i+1]
			}
		}

		if !change {
			break
		}
	}

	for i := 0; i < KEEP_BEST_COUNT; i++ {
		buffer = append(buffer, cornerTrainer.Params[i])
	}

	for i := 0; i < 5; i++ {
		buffer = append(buffer, cornerTrainer.Mutate(cornerTrainer.Params[i], MUTATION_THRESHOLD))
		buffer = append(buffer, cornerTrainer.Mutate(cornerTrainer.Params[i], MUTATION_THRESHOLD*2))
	}

	for i := 0; i < 5; i++ {
		buffer = append(buffer, cornerTrainer.PairParams(cornerTrainer.Params[0], cornerTrainer.Params[i+1]))
	}

	for i := 0; i < 3; i++ {
		buffer = append(buffer, cornerTrainer.PairParams(cornerTrainer.Params[1], cornerTrainer.Params[i+2]))
	}

	for i := 0; i < 2; i++ {
		buffer = append(buffer, cornerTrainer.PairParams(cornerTrainer.Params[2], cornerTrainer.Params[i+3]))
	}

	cornerTrainer.Params = buffer
}

func (cornerTrainer *CornerTrainer) PairParams(param1 CornerParams, param2 CornerParams) CornerParams {
	child := CornerParams{}
	child.NeuralNet = make([]float32, cornerTrainer.NeuralNet.TotalConnectionsCount())

	childVal := reflect.ValueOf(&child)
	param1Val := reflect.ValueOf(param1)
	param2Val := reflect.ValueOf(param2)

	for i := 1; i < reflect.TypeOf(child).NumField(); i++ {
		value := childVal.Elem().Field(i)

		switch value.Kind() {
		case reflect.Int:
			newValue := (param1Val.Field(i).Int() + param2Val.Field(i).Int()) / 2
			value.SetInt(newValue)
		case reflect.Float32:
			newValue := (param1Val.Field(i).Float() + param2Val.Field(i).Float()) / 2
			value.SetFloat(newValue)
		case reflect.Float64:
			newValue := (param1Val.Field(i).Float() + param2Val.Field(i).Float()) / 2
			value.SetFloat(newValue)
		case reflect.Slice:
			for j := 0; j < cornerTrainer.NeuralNet.TotalConnectionsCount(); j++ {
				value.Index(j).SetFloat((param1Val.Field(i).Index(j).Float() + param2Val.Field(i).Index(j).Float()) / 2)
			}
		default:
			panic("Unspecified behaviour in PairParams")
		}

	}

	return child
}

func (cornerTrainer *CornerTrainer) Mutate(params CornerParams, mutationThreshold float32) CornerParams {
	values := reflect.ValueOf(&params)
	limits := reflect.ValueOf(cornerTrainer.ParamsLimit)

	for i := 1; i < reflect.TypeOf(values).NumField(); i++ {
		value := values.Elem().Field(i)

		if rand.Float32() > mutationThreshold {
			value.Set(limits.Field(i - 1).MethodByName("Mutate").Call([]reflect.Value{value.Addr()})[0])
		}
	}

	return params
}
