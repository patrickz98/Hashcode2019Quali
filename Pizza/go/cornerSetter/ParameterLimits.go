package cornerSetter

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

const MUTATE_FACTOR = 5

type CornerParamsLimit struct {
	NeuralNet   LimitFloat32NNSlice
	InputOrder  LimitFloat32NNSlice
	InputTarget LimitIntSlice
}

type LimitInt struct {
	Min int
	Max int
}

type LimitFloat32 struct {
	Min float32
	Max float32
}

type LimitIntSlice struct {
	Count    int
	MaxInt   int
	Mutating int
}

type LimitFloat32NNSlice struct {
	Count    int
	Mutating int
}

func (lim LimitInt) GetRandomNumber() int {
	return lim.Min + rand.Intn(lim.Max-lim.Min)
}

func (lim LimitFloat32) GetRandomNumber() float32 {
	return lim.Min + (rand.Float32() * (lim.Max - lim.Min))
}

func (lim LimitFloat32NNSlice) GetRandomNumber() []float32 {
	buffer := make([]float32, lim.Count)

	for i := 0; i < lim.Count; i++ {
		buffer[i] = (rand.Float32() * 2) - 1
	}

	return buffer
}

func (lim LimitIntSlice) GetRandomNumber() []int {
	buffer := make([]int, lim.Count)

	for i := 0; i < lim.Count; i++ {
		buffer[i] = rand.Intn(lim.MaxInt)
	}

	return buffer
}

func (lim LimitInt) Mutate(inp *int64) int64 {
	return int64(MinMax(float32(lim.Max), float32(lim.Min), float32(*inp)+float32(rand.NormFloat64())*float32(lim.Max-lim.Min)/MUTATE_FACTOR))
}

func (lim LimitFloat32) Mutate(inp *float32) float32 {
	return MinMax(lim.Min, lim.Max, *inp+float32(rand.NormFloat64()*float64(lim.Max-lim.Min)/MUTATE_FACTOR))
}

func (lim LimitFloat32NNSlice) Mutate(data *[]float32) []float32 {

	for i := 0; i < rand.Intn(lim.Mutating); i++ {
		index := int(rand.Float32() * float32(lim.Count))

		(*data)[index] = MinMax(-1, 1, (*data)[index]+float32(rand.NormFloat64()*float64(2)/MUTATE_FACTOR))
	}

	return *data
}

func (lim LimitIntSlice) Mutate(data *[]int) []int {

	for i := 0; i < rand.Intn(lim.Mutating); i++ {
		index := int(rand.Float32() * float32(lim.Count))

		(*data)[index] = rand.Intn(lim.MaxInt)
	}

	return *data
}

func (cpl CornerParamsLimit) getRandomParams() CornerParams {
	newParams := CornerParams{}

	values := reflect.ValueOf(&newParams)
	limits := reflect.ValueOf(cpl)

	for i := 1; i < reflect.TypeOf(newParams).NumField(); i++ {
		value := values.Elem().Field(i)

		value.Set(limits.Field(i - 1).MethodByName("GetRandomNumber").Call([]reflect.Value{})[0])

		switch value.Kind() {
		case reflect.String:
			newParams.Name += value.String()
		case reflect.Int:
			newParams.Name += strconv.FormatInt(value.Int(), 10)
		case reflect.Float32:
			newParams.Name += fmt.Sprintf("%f", value.Float())
		default:
		}
	}

	return newParams
}
