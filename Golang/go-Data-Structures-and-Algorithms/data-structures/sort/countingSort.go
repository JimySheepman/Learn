package main

import "fmt"

func main() {
	arr := []int32{1, 2, 1, 1, 3}
	countArr := make([]int32, len(arr))
	output := make([]int32, len(arr))
	dict := make(map[int32]int32)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	for i, v := range dict {
		countArr[i] = v
	}
	for i := 0; i < len(countArr)-1; i++ {
		countArr[i+1] = countArr[i] + countArr[i+1]
	}
	fmt.Println(arr)
	fmt.Println(countArr)
	for _, v := range arr {
		output[countArr[v]-1] = v
		countArr[v] -= 1
	}
	fmt.Println(output)
}
