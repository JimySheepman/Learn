package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	positiveCount := 0
	negativeCount := 0
	for _, v := range numbers {
		if v < 0 {
			negativeCount += v
		} else if v == 0 {
			continue
		} else {
			positiveCount++
		}
	}
	res = append(res, positiveCount, negativeCount)
	return res // your code here
}
