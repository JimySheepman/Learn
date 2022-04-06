package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(4)

	fmt.Println(r.Len())
}
