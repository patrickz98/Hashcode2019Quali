package cornerSetter

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

const MUTATE_FACTOR = 20

type CornerParamsLimit struct {
	CornerMinScore LimitFloat32
	NeuralNet      LimitFloat32NNSlice
}

type LimitInt struct {
	Min int
	Max int
}

type LimitFloat32 struct {
	Min float32
	Max float32
}

type LimitFloat32NNSlice struct {
	NeuralNetConnections int
}

func (lim LimitInt) GetRandomNumber() int {
	return lim.Min + rand.Intn(lim.Max-lim.Min)
}

func (lim LimitFloat32) GetRandomNumber() float32 {
	return lim.Min + (rand.Float32() * (lim.Max - lim.Min))
}

func (lim LimitFloat32NNSlice) GetRandomNumber() []float32 {
	buffer := make([]float32, 0)

	for i := 0; i < lim.NeuralNetConnections; i++ {
		buffer = append(buffer, rand.Float32())
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

	for i := 0; i < rand.Intn(10); i++ {
		index := int(rand.Float32() * float32(lim.NeuralNetConnections))

		(*data)[index] = MinMax(0, 1, (*data)[index]+float32(rand.NormFloat64()*float64(1)/MUTATE_FACTOR))
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
