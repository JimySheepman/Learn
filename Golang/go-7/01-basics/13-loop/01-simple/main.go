package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i+1)
	}

	fmt.Println()

	j := 0
	for j < 10 {
		fmt.Printf("%d", j+1)
		j++
	}
}
