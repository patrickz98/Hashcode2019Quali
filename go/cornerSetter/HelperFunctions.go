package cornerSetter

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinMax(min float32, max float32, inp float32) float32 {
	if inp > max {
		return max
	}

	if inp < min {
		return min
	}

	return inp
}
