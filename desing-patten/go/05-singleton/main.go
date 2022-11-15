package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()

		go getOnceInstance()
	}

	fmt.Scanln()
}
