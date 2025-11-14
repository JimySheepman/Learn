package kata

func multipleOfIndex(ints []int) []int {
	res := make([]int, 0)

	for i, v := range ints {
		if i > 0 && v%i == 0 {
			res = append(res, v)
		}
	}
	return res
}
