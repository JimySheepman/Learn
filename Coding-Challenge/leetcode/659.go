func isPossible(nums []int) bool {
	var list [][]int
LOOP:
	for _, n := range nums {
		for i := len(list) - 1; i >= 0; i-- {
			if list[i][len(list[i])-1]+1 == n {
				list[i] = append(list[i], n)
				continue LOOP
			}
		}
		list = append(list, []int{n})
	}

	for _, v := range list {
		if len(v) < 3 {
			return false
		}
	}
	return true
}