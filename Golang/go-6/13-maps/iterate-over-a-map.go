package main

import "fmt"

func main() {
	var employee = map[string]int{"mark": 10, "sandy": 20, "rocky": 30, "rajiv": 40, "kate": 50}
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

}
