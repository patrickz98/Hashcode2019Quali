package cornerSetter

import "reflect"

type CornerParams struct {
	Name        string
	NeuralNet   []float32
	InputOrder  []float32
	InputTarget []int
}

func (cornerParams *CornerParams) Init(data []string) {
	values := reflect.ValueOf(&cornerParams)

	for i := 1; i < reflect.TypeOf(cornerParams).NumField(); i++ {
		value := values.Elem().Field(i)

		value.Set(reflect.ValueOf(data[i]))
	}
}
