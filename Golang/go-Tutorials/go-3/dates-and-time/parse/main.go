package main

import (
	"fmt"
	"time"
)

func main() {
	timeToParse := "2019-02-15T07:33-05:00"
	t, err := time.Parse("2006-01-02T15:04-07:00", timeToParse)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
}
