package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)

	normalHouse := director.buildHouse()
	printDetails(normalHouse, "Normal")

	director.setBuilder(iglooBuilder)

	iglooHouse := director.buildHouse()
	printDetails(iglooHouse, "Igloo")
}

func printDetails(h House, s string) {
	fmt.Printf("%s House Door Type: %s\n", s, h.doorType)
	fmt.Printf("%s House Window Type: %s\n", s, h.windowType)
	fmt.Printf("%s House Num Floor: %d\n", s, h.floor)
}
