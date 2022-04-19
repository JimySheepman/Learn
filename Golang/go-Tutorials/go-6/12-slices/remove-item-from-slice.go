package main

import "fmt"

func main() {
	var strSlice = []string{"India", "Canda", "Japan", "Germnay", "Italy"}
	fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 3)
	fmt.Println(strSlice)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
