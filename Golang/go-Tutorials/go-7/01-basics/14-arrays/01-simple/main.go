package main

import "fmt"

func main() {

	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("numbers: ", numbers)

	var items [6]string
	items[0] = "G"
	items[1] = "O"
	items[2] = "P"
	items[3] = "H"
	items[4] = "E"
	items[5] = "R"
	fmt.Println("items:", items)

	values := [...]int{1, 2, 3, 4}
	fmt.Println("values:", values)

	names := [4]string{
		"Go",
		"Gopher",
	}
	fmt.Println(names)

	fmt.Printf("names length: %d\n", len(names))
}
