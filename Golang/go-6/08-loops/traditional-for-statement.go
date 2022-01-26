package main

import "fmt"

func main() {
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}
