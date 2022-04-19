package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("file:", file)
	}
	myError := errors.New("Bu bir hata!")
	fmt.Println(myError.Error())

	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54
	fmt.Println(myArray1)
	fmt.Println(myArray1[2])
	var colors [3]string
	colors[0] = "red"
	colors[1] = "yellow"
	colors[2] = "blue"
	fmt.Println(colors)

	myArray2 := [...]int{4, 3, 1, 5, 6, 7, 8, 123} //  auto size
	fmt.Println(myArray2)
	fmt.Println(len(myArray2), cap(myArray2))

	fmt.Println("-----")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		log.Fatal("fatal error:", err)
		os.Exit(1)
	}
}
