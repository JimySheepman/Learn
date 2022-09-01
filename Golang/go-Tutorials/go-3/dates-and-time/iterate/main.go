package main

import (
	"fmt"
	"time"
)

func main() {
	start, err := time.Parse("2006-01-02", "2019-02-19")
	if err != nil {
		panic(err)
	}
	end, err := time.Parse("2006-01-02", "2020-07-17")
	if err != nil {
		panic(err)
	}
	for i := start; i.Unix() < end.Unix(); i = i.AddDate(0, 0, 1) {
		fmt.Println(i.Format(time.RFC3339))
	}
	for i := start; i.Unix() < end.Unix(); i = i.Add(time.Minute * 2) {
		fmt.Println(i.Format(time.RFC3339))
	}
}
