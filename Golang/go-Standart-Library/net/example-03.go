package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	var d net.Dialer
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()

	d.LocalAddr = nil
	raddr := net.UnixAddr{Name: "/path/to/unix.sock", Net: "unix"}
	conn, err := d.DialContext(ctx, "unix", raddr.String())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, socket!")); err != nil {
		log.Fatal(err)
	}
}
