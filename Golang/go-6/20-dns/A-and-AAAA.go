package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupIP("facebook.com")
	for _, v := range iprecords {
		fmt.Println(v)
	}
}
