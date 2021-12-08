package main

import "fmt"

func main() {

	var m1 = make(map[string]int)

	m1["John"] = 1
	m1["Jane"] = 2

	fmt.Printf("m1: %s\n", m1)
	fmt.Println(m1["Jane"])

	m2 := map[string]int{
		"Jimmy":   1,
		"Merlins": 2,
	}
	fmt.Println(m2["Jimmy"])
	fmt.Println(m2["Merlins"])

	m3 := map[string]int{
		"John": 1,
		"Jane": 2,
		"Jack": 3,
		"Jill": 4,
	}
	fmt.Printf("Before deleted entry:\n")
	for x := range m3 {
		fmt.Println(m3[x])
	}

	delete(m3, "Jack")
	fmt.Printf("\nAfter deleted entry:\n")
	for x := range m3 {
		fmt.Println(m3[x])
	}
}
