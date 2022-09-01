package main

import "fmt"

func main() {
	languages := []string{"Java", "PHP", "C"}
	fmt.Println("Capacity :", cap(languages))

	addGo(languages)

	fmt.Println("Capacity :", cap(languages))
	fmt.Println(languages)
}

func addGo(languages []string) {
	languages = append(languages, "Go")
	fmt.Println("in function, capacity", cap(languages))
}
