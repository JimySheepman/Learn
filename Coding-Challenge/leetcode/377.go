func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 1
	return helper(nums, target, &dp)
}

func helper(nums []int, target int, dp *[]int) int {

	if (*dp)[target] != -1{
		return (*dp)[target]
	}
	var res int
	for i := 0; i<len(nums); i++ {
		if target >= nums[i] {
			res += helper(nums, target-nums[i], dp)
		}
	}
	(*dp)[target] = res
	return res
}

