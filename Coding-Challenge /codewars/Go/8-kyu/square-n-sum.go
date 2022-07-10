package kata

func SquareSum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v * v
	}
	return sum
}
