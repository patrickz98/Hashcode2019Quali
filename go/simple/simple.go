package simple

import (
	"encoding/json"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrettyJson(data interface{}) (bytes []byte, err error) {
	return json.MarshalIndent(data, "", "    ")
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
