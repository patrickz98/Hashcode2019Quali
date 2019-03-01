package simple

import (
	"encoding/json"
	"fmt"
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

func FindRecursive(set []int, targetSum int, currentSum int, currentIndex int) [][]int {

	results := make([][]int, 0)

	for i := currentIndex; i < len(set); i++ {

		newSum := currentSum + set[ i ]
		if newSum > targetSum {
			continue
		}

		if newSum == targetSum {
			results = append(results, []int{ set[i] })
		}

		result := FindRecursive(set, targetSum, newSum, i + 1)
		if result == nil {
			continue
		}

		for inx := range result {
			result[ inx ] = append(result[ inx ], set[ i ])
			results = append(results, result[ inx ])
		}
	}

	return results
}

func FindDP(set []int, sum int) [][]int {

	sets := make([][]int, 0)
	solution := make([][]bool, len(set) + 1)

	for inx := range solution {
		solution[ inx ] = make([]bool, sum + 1)
		solution[ inx ][ 0 ] = true
	}

	for iny := 1; iny <= len(set); iny++ {

		for inx := 1; inx <= sum; inx++ {

			solution[ iny ][ inx ] = solution[ iny - 1 ][ inx ]

			if !solution[ iny ][ inx ] && inx >= set[ iny - 1] {
				// solution[ iny ][ inx ] = solution[ iny ][ inx ] || solution[ iny - 1 ][ inx - set[ iny - 1 ]]
				solution[ iny ][ inx ] = solution[ iny - 1 ][ inx - set[ iny - 1 ]]

				// fmt.Printf("(%d, %d) --> %d, %d >> %v\n", iny, inx, iny - 1, inx - set[ iny - 1 ], solution[ iny ][ inx ])
			}
		}

		if !solution[ iny ][ sum ] {
			continue
		}

		result := make([]int, 0)
		q := sum

		for p := iny - 1; p >= 0; p-- {

			if solution[ p ][ q ] {
				continue
			}

			s := set[ p ]
			result = append(result, s)
			q -= s
		}

		sets = append(sets, result)
	}

	for _, sol := range solution {
		fmt.Println(sol)
	}

	return sets
}

func Test() {

	sum := 8
	set := []int{ 1, 2, 5, 7 }
	// set := []int{ 7, 2, 5, 1 }

	result := FindRecursive(set, sum, 0, 0)

	fmt.Println(len(result))
	fmt.Println(result)

	Exit()
}
