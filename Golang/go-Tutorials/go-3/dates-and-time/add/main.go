package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	now = now.AddDate(12, 1, 3)

	now = now.Add(time.Nanosecond * 1)
	now = now.Add(time.Microsecond * 5)
	now = now.Add(time.Millisecond * 5)
	now = now.Add(time.Second * 5)
	now = now.Add(time.Minute * 5)
	now = now.Add(time.Hour * 5)

	fmt.Println(now)
}
