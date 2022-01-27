package main

import "fmt"

func main() {
	var strSlice = []string{"India", "Canada", "Japan"}
	fmt.Println(strSlice)

	strSlice[2] = "Germany"
	fmt.Println(strSlice)
}
