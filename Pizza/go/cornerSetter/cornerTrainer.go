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

const POPULATION = 50
const KEEP_BEST = 0.20
const MUTATE = 0.50
const PAIR = 0.30

const MUTATION_THRESHOLD = 0

var MaxSlicesInAPlace = 0

type CornerTrainer struct {
	Slicer             *slicer.Slicer
	ParamsLimit        CornerParamsLimit
	Generation         int
	LastImprovementGen int
	HighScore          float32
	OldHighScore       float32
	NeuralNet          NeuralNet

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
	cornerTrainer.NeuralNet.UseJFMethod = false

	cornerTrainer.CornerSetter = CornerSetter{
		Slicer:    cornerTrainer.Slicer,
		Pizza:     cornerTrainer.Slicer.Pizza,
		Params:    CornerParams{},
		NeuralNet: &cornerTrainer.NeuralNet,
	}

	cornerTrainer.buildSlicesCache()

	cornerTrainer.Params = make([]CornerParams, 0)
	cornerTrainer.Scores = make([]float32, 0)

	cornerTrainer.ParamsLimit = CornerParamsLimit{
		LimitFloat32NNSlice{cornerTrainer.NeuralNet.TotalConnectionsCount(), 3},
		LimitFloat32NNSlice{len(cornerTrainer.NeuralNet.InputLayer), 1},
		LimitIntSlice{len(cornerTrainer.NeuralNet.InputLayer), len(cornerTrainer.NeuralNet.Layer[0]), 1},
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
	averageKeeperScore := float32(0)
	oldAverageKeeperScore := float32(0)
	KeeperAverageImprovementGen := 0
	cornerTrainer.Scores = make([]float32, POPULATION)
	cornerTrainer.LastImprovementGen = 0
	for cornerTrainer.HighScore < 1 {
		scoreSum := float32(0)
		scoreKeeperSum := float32(0)

		fmt.Printf("\nCurrent Generation: %d, AverageScore %.3f/%.3f, Highscore: %.3f, Improvement: %.3f/%d, %.3f/%d\n", cornerTrainer.Generation, averageKeeperScore, averageScore, 100*cornerTrainer.HighScore, averageKeeperScore-oldAverageKeeperScore, KeeperAverageImprovementGen, 100*(cornerTrainer.HighScore-cornerTrainer.OldHighScore), cornerTrainer.LastImprovementGen)
		for i, params := range cornerTrainer.Params {

			addToScore := cornerTrainer.Scores[i]

			Rep := 0
			if i >= int(KEEP_BEST*POPULATION) {
				cornerTrainer.CornerSetter.Params = params

				cornerTrainer.CornerSetter.Pizza.RemoveAllSlice()

				for j := 1; j <= 5; j++ {

					if cornerTrainer.CornerSetter.SetSlices(j, 1-(0.2*float32(j))) {
						Rep++
					}

					/*if !cornerTrainer.CornerSetter.SetSlices(j, 1 - (0.1 * float32(j))) {
						break
					}*/
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

			if i < int(KEEP_BEST*POPULATION) {
				scoreKeeperSum += addToScore
			}

			fmt.Printf("%d: %.3f, %d\n", i, addToScore*100, Rep)
			cornerTrainer.Scores[i] = addToScore
			//fmt.Printf("Current Generation: %d/%d, Score %.3f, AverageScore %.3f, Highscore: %.3f, Improvement: %.3f/%d, HighestRep %d,             \r", i, cornerTrainer.Generation, 100*addToScore, averageScore, 100*cornerTrainer.HighScore, 100*(cornerTrainer.HighScore-cornerTrainer.OldHighScore), LastImprovementGen, HighestRep)
		}
		averageScore = 100 * scoreSum / float32(len(cornerTrainer.Params))
		averageKeeperScore = 100 * scoreKeeperSum / float32(int(KEEP_BEST*POPULATION))

		if averageKeeperScore > oldAverageKeeperScore {
			oldAverageKeeperScore = averageKeeperScore
			KeeperAverageImprovementGen = cornerTrainer.Generation
		}

		for j, _ := range cornerTrainer.Params {
			if cornerTrainer.Scores[j] > cornerTrainer.HighScore {

				cornerTrainer.OldHighScore = cornerTrainer.HighScore
				cornerTrainer.HighScore = cornerTrainer.Scores[j]
				cornerTrainer.LastImprovementGen = cornerTrainer.Generation
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

	for i := 0; i < int(POPULATION*KEEP_BEST); i++ {
		buffer = append(buffer, cornerTrainer.Params[i])
	}

	for i := 0; i < int(POPULATION*MUTATE); i++ {
		buffer = append(buffer, cornerTrainer.Mutate(cornerTrainer.Params[i], MUTATION_THRESHOLD))
	}

	for i := 0; i < int(POPULATION*PAIR); i += 2 {
		buffer = append(buffer, cornerTrainer.Mutate(cornerTrainer.PairParams(cornerTrainer.Params[rand.Intn(int(KEEP_BEST*POPULATION))], cornerTrainer.Params[rand.Intn(int(KEEP_BEST*POPULATION))]), 0))
		buffer = append(buffer, cornerTrainer.PairParams(cornerTrainer.Params[rand.Intn(int(KEEP_BEST*POPULATION))], cornerTrainer.ParamsLimit.getRandomParams()))
	}

	for len(buffer) < POPULATION {
		buffer = append(buffer, cornerTrainer.ParamsLimit.getRandomParams())
	}

	buffer = buffer[:POPULATION]

	cornerTrainer.Params = buffer
}

func (cornerTrainer *CornerTrainer) PairParams(param1 CornerParams, param2 CornerParams) CornerParams {
	child := CornerParams{}
	child.NeuralNet = make([]float32, cornerTrainer.NeuralNet.TotalConnectionsCount())
	child.InputOrder = make([]float32, len(cornerTrainer.NeuralNet.InputLayer))
	child.InputTarget = make([]int, len(cornerTrainer.NeuralNet.InputLayer))

	childVal := reflect.ValueOf(&child)
	param1Val := reflect.ValueOf(param1)
	param2Val := reflect.ValueOf(param2)

	for i := 1; i < reflect.TypeOf(child).NumField(); i++ {
		value := childVal.Elem().Field(i)

		switch value.Kind() {
		case reflect.Int:
			//newValue := (param1Val.Field(i).Int() + param2Val.Field(i).Int()) / 2
			//value.SetInt(newValue)

			if rand.Float32() > 0.5 {
				value.SetInt(param1Val.Field(i).Int())
			} else {
				value.SetInt(param2Val.Field(i).Int())
			}

		case reflect.Float32:
			//newValue := (param1Val.Field(i).Float() + param2Val.Field(i).Float()) / 2
			//value.SetFloat(newValue)

			if rand.Float32() > 0.5 {
				value.SetFloat(param1Val.Field(i).Float())
			} else {
				value.SetFloat(param2Val.Field(i).Float())
			}
		case reflect.Float64:
			//newValue := (param1Val.Field(i).Float() + param2Val.Field(i).Float()) / 2
			//value.SetFloat(newValue)

			if rand.Float32() > 0.5 {
				value.SetFloat(param1Val.Field(i).Float())
			} else {
				value.SetFloat(param2Val.Field(i).Float())
			}

		case reflect.Slice:
			for j := 0; j < int(reflect.ValueOf(cornerTrainer.ParamsLimit).Field(i-1).FieldByName("Count").Int()); j++ {
				//value.Index(j).SetFloat((param1Val.Field(i).Index(j).Float() + param2Val.Field(i).Index(j).Float()) / float64(2))

				if rand.Float32() > 0.5 {
					value.Index(j).Set(param1Val.Field(i).Index(j))
				} else {
					value.Index(j).Set(param2Val.Field(i).Index(j))
				}

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
