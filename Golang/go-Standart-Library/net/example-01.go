package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.CIDRMask(31, 32))
	fmt.Println(net.CIDRMask(64, 128))
}
