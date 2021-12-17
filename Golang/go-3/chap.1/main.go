package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	str1 := "the quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	stringLength, _ := fmt.Println(str1, str2, str3)
	fmt.Println("String length:", stringLength)
	fmt.Printf("Value of aNumber : %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", true)
	fmt.Printf("Valu of aNumber as float: %.2f\n", float32(aNumber))

	fmt.Printf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, true)
	myString := fmt.Sprintf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, true)
	fmt.Print(myString)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", f)
	}

	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Go Launch at %s\n", t)

	fmt.Println("-----")

	now := time.Now()
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("-----")

	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())

	fmt.Println("-----")

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf(
		"Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	fmt.Println("-----")

	longFormat := "Monday, Junuary 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	fmt.Printf("-----")

	shortFormat := "1/2/06"
	fmt.Printf("Tomorrow is ", tomorrow.Format(shortFormat))

	for i := 0; i < 5; i++ {
		fmt.Println("Value: ", i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println("sum: ", sum)
	}

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("index: %d , value: %d\n", i, v)
	}

	capitalCity := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}
	for key, value := range capitalCity {
		fmt.Printf("Key %s : Value %s \n", key, value)
	}
}
