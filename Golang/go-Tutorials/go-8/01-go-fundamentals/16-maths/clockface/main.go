package main

import (
	"os"
	"time"

	"math/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
