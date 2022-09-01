package main

import (
	"fmt"
	"time"
)

func main() {

	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	firstJanuary1980 := time.Date(1980, 1, 1, 0, 0, 0, 0, location)

	timeToParse := "2019-02-15T07:33-02:00"
	t, err := time.Parse("2006-01-02T15:04-07:00", timeToParse)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	if t.After(firstJanuary1980) && t.Before(now) {
		fmt.Println(t, "is between ", firstJanuary1980, "and", now)
	} else {
		fmt.Println("not in between")
	}
}
