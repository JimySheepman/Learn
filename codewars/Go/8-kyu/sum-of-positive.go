package kata

func PositiveSum(numbers []int) int {
	result := 0
	for _, i := range numbers {
		if i > 0 {
			result += i
		}
	}
	return result
}
