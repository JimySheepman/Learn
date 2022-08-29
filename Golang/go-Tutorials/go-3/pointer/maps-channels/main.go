package main

import "log"

func addElement(cities map[string]string) {
	cities["France"] = "Paris"
}

func main() {
	cities := make(map[string]string)
	addElement(cities)
	log.Println(cities)
}
