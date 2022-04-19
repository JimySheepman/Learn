package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

func main() {
	fmt.Println("BeginningOfMinute \t", now.BeginningOfMinute())
	fmt.Println("BeginningOfHour \t", now.BeginningOfHour())
	fmt.Println("BeginningOfDay  \t", now.BeginningOfDay())
	fmt.Println("BeginningOfWeek \t", now.BeginningOfWeek())
	fmt.Println("BeginningOfMonth \t", now.BeginningOfMonth())
	fmt.Println("BeginningOfQuarter \t", now.BeginningOfQuarter())
	fmt.Println("BeginningOfYear \t", now.BeginningOfYear())
	fmt.Println()

	fmt.Println("EndOfMinute \t", now.EndOfMinute())
	fmt.Println("EndOfHour \t", now.EndOfHour())
	fmt.Println("EndOfDay  \t", now.EndOfDay())
	fmt.Println("EndOfWeek \t", now.EndOfWeek())
	fmt.Println("EndOfMonth \t", now.EndOfMonth())
	fmt.Println("EndOfQuarter \t", now.EndOfQuarter())
	fmt.Println("EndOfYear \t", now.EndOfYear())
	fmt.Println()

	fmt.Println("Monday 	\t", now.Monday())
	fmt.Println("Sunday 	\t", now.Sunday())
	fmt.Println("EndOfSunday \t", now.EndOfSunday())
	fmt.Println()

	fmt.Println(now.Parse("2017"))
	fmt.Println(now.Parse("2017-12-12 12:20"))

	t := time.Date(2020, 07, 18, 17, 51, 49, 123456789, time.Now().Location())
	fmt.Println(now.With(t).EndOfMonth())
}
