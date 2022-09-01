package main

import "fmt"

func printer() func() {
	k := 1
	return func() {
		fmt.Printf("Print n. %d\n", k)
		k++
	}
}

func main() {

	p := printer()
	p()
	p()
	p()
}
