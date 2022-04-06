package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}
}
