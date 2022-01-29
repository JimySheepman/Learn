package main

import "fmt"

func main() {
	var rad float64
	const PI float64 = 3.14
	var area float64
	var ci float64

	fmt.Print("Enter radius of Circle:")
	fmt.Scanln(&rad)
	area = PI * rad * rad
	fmt.Println("Area of Circle is:", area)
	ci = 2 * PI * rad
	fmt.Println("Circumference of Circle is: ", ci)
}
