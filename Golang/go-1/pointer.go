package main

import "fmt"

func main() {
	num := 3
	fmt.Println("Variable value: ", num)
	fmt.Println("Variable address: ", &num)
	num2 := 3
	numAD := &num2
	fmt.Println("Variable value:", num)
	fmt.Println("Variable address:", numAD)  // address
	fmt.Println("Value at address:", *numAD) // address value

	var num3 int = 8
	var ptr *int
	ptr = &num3
	fmt.Println("ptr (adress of num):", ptr)
	fmt.Println("ptr (value of num):", *ptr)

}
