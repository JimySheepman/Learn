func singleNumber(nums []int) int {
	uniqNum := 0
	for _, num := range nums {
		uniqNum ^= num
	}
	return uniqNum
}
