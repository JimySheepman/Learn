package main

import "fmt"

func main() {
	customers := [4]string{"John Doe", "Helmuth Verein", "Dany Beril", "Oliver Lump"}
	customersSlice := customers[0:1]
	fmt.Println(customersSlice)
	customers[0] = "John Doe Modified"
	fmt.Println("After modification of original array")
	fmt.Println(customersSlice)
}
