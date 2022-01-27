package main

import (
	"fmt"
	"net"
)

func main() {
	ptr, _ := net.LookupAddr("6.8.8.8")
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}
}
