package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	results := make([]int, len(arr))

	for i, x := range arr {
		sqrt := math.Sqrt(float64(x))

		if sqrt == math.Floor(sqrt) {
			results[i] = int(sqrt)
		} else {
			results[i] = x * x
		}
	}

	return results
}
