package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%s\n", now)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	nowNYC := now.In(loc)
	fmt.Printf("%s\n", nowNYC)
}
