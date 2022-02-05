package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	Sample("IPv4", IPv4)
	Sample("LookupIP", LookupIP)
	Sample("ParseIP", ParseIP)
	Sample("Listener", Listener)
	Sample("Dialer", Dialer)
}

func IPv4() {
	fmt.Println(net.IPv4(8, 8, 8, 8))
}

func LookupIP() {
	ips, err := net.LookupIP("golang.org")
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func ParseIP() {
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(net.ParseIP("2001:db8::68"))
	fmt.Println(net.ParseIP("192.0.2"))
}

func Listener() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}

func Dialer() {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", ":2000")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		fmt.Println("error:", err)
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
