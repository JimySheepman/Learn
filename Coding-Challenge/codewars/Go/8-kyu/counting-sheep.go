package kata

func CountSheeps(numbers []bool) int {
	count := 0
	for _, j := range numbers {
		if j == true {
			count++
		}

	}
	return count
}
