package main

import "fmt"

func MaxFrequencyNumber(array []int64) int64 {
	frequencyMap := map[int64]int64{}

	for _, value := range array {
		frequencyMap[value] += 1
	}

	result := FindMax(frequencyMap)
	return result
}

func FindMax(frequencyMap map[int64]int64) int64 {
	var max int64 = 0
	var index int64 = 0

	for i, v := range frequencyMap {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func main() {
	s := MaxFrequencyNumber([]int64{4, 4, 4, 3, 4, 4, 4, 4, 2, 2, 2, 1, 1, 5, 5})
	fmt.Println(s)
}
