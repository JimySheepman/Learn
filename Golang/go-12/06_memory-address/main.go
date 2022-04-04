package main

import "fmt"

func showing_address() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
}

const metersToYards float64 = 1.09361

func using_address() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yars.")

}

func main() {
	showing_address()
	using_address()
}
