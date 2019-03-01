package simple

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrettyJson(data interface{}) (bytes []byte, err error) {
	return json.MarshalIndent(data, "", "    ")
}

func PrettyJsonString(data interface{}) (str string, err error) {

	bytes, err := PrettyJson(data)
	return string(bytes), err
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(values ...int) int {

	if len(values) <= 0 {
		return 0
	}

	smallest := values[ 0 ]

	for _, val := range values {

		if smallest > val {
			smallest = val
		}
	}

	return smallest
}

func Exit() {
	os.Exit(0)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Write(path string, data string) {

	bytes := []byte(data)
	err := ioutil.WriteFile(path, bytes, 0644)
	CheckErr(err)
}
