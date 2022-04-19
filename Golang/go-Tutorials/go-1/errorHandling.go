package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Add(6, -3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	defer consoleLog2("Hello")
	consoleLog2("there")

	result3 := div(5, 0)
	fmt.Println(result3)

	result3 = div(5, 2)
	fmt.Println(result3)

	result2 := Add2(5, -3)
	fmt.Println(result2)

}

func Add(num1 int, num2 int) (int, error) {
	err := errors.New("Number cannot be negative")
	if num1 < 0 || num2 < 0 {
		return 0, err
	}
	return num1 + num2, nil
}
func Add2(num1 int, num2 int) int {

	if num1 < 0 || num2 < 0 {
		panic("Number cannot be negative")
	}

	return num1 + num2
}
func consoleLog2(msg string) {

	fmt.Println(msg)
}
func div(num1 float64, num2 float64) float64 {

	defer func() {

		fmt.Println(recover())
	}()

	return num1 / num2
}
