package main

import (
	"errors"
	"fmt"
	"testing/iotest"
)

func main() {
	r := iotest.ErrReader(errors.New("custom error"))
	n, err := r.Read(nil)
	fmt.Printf("n:   %d\nerr: %q\n", n, err)
}
