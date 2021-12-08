package main

import "fmt"
import "time"

func main() {

	go consoleLog3("World")
	go consoleLog3("Hello")

	time.Sleep(100 * time.Millisecond)
	fmt.Println(getNum())
}

func consoleLog3(msg string) {

	fmt.Println(msg)
}
func getNum() int {

	i := -1

	go func() {

		i = 5
	}()

	return i
}
