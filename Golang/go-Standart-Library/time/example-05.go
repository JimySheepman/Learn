package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("One second is %d milliseconds.\n", u.Milliseconds())
}
