package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp", "192.0.2.1:2000")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.WriteTo([]byte("data"), dst)
	if err != nil {
		log.Fatal(err)
	}
}
