package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("port", 4242, "the port on which the server will listen")
	flag.Parse()
	fmt.Printf("%d", *port)

	var port2 int
	flag.IntVar(&port2, "port2", 4242, "the port on which the server will listen")
	flag.Parse()
	fmt.Printf("%d\n", port2)
}
