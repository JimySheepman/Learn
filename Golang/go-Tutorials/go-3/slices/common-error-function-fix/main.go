package main

import "fmt"

func main() {
	languages := []string{"Java", "PHP", "C"}
	fmt.Println("Capacity :", cap(languages))

	addGoFixed(&languages)

	fmt.Println("Capacity :", cap(languages))
	fmt.Println(languages)
}

func addGoFixed(languages *[]string) {
	*languages = append(*languages, "Go")
}
