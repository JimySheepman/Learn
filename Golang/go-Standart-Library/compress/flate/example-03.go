package main

import (
	"compress/flate"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	rp, wp := io.Pipe()

	wg.Add(1)
	go func() {
		defer wg.Done()

		zw, err := flate.NewWriter(wp, flate.BestSpeed)
		if err != nil {
			log.Fatal(err)
		}

		b := make([]byte, 256)
		for _, m := range strings.Fields("A long time ago in a galaxy far, far away...") {
			b[0] = uint8(copy(b[1:], m))

			if _, err := zw.Write(b[:1+len(m)]); err != nil {
				log.Fatal(err)
			}
			if err := zw.Flush(); err != nil {
				log.Fatal(err)
			}
		}
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		zr := flate.NewReader(rp)
		b := make([]byte, 256)
		for {
			if _, err := io.ReadFull(zr, b[:1]); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			n := int(b[0])
			if _, err := io.ReadFull(zr, b[:n]); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Recived %d bytes: %s\n", n, b[:n])
		}
		fmt.Println()

		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
