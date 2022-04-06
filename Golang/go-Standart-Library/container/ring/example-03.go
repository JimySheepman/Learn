package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(2)
	s := ring.New(2)

	lr := r.Len()
	ls := r.Len()

	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}

	rs := r.Link(s)
	rs.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})
}
