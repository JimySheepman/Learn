package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	for i, num := range nums {
		nums[i-count] = nums[i]
		if num == val {
			count += 1
		}
	}
	return len(nums) - count
}

func main() {
	// nums := []int{3, 2, 2, 3}
	//val := 3
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
