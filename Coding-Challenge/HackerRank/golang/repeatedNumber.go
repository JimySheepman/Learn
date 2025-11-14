package main

import "fmt"

func RepeatedNumber(array []int64) map[int64]int64 {
	frequencyMap := map[int64]int64{}

	for _, value := range array {
		frequencyMap[value] += 1
	}

	return frequencyMap
}

func main() {
	s := RepeatedNumber([]int64{4, 4, 4, 3, 4, 4, 4, 4, 2, 2, 2, 1, 1, 5, 5})
	fmt.Println(s)
}
