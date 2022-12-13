package main

import "fmt"

// Client
//
// The Client must associate one of the builder objects with the director. Usually,
// it’s done just once, via parameters of the director’s constructor. Then the director
// uses that builder object for all further construction.  However, there’s an alternative
// approach for when the client passes the builder object to the production method of the director.
// In this case, you can use a different builder each time you produce something with the director.
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
