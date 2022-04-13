package main

import (
	"fmt"
	"net"
)

func main() {
	ipv4Addr := net.ParseIP("192.0.2.1")
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Addr.Mask(ipv4Mask))

	ipv6Addr := net.ParseIP("2001:db8:a0b:12f0::1")
	ipv6Mask := net.CIDRMask(32, 128)
	fmt.Println(ipv6Addr.Mask(ipv6Mask))
}
