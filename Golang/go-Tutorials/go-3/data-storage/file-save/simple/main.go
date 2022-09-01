package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("data-storage/file-save/simple/test.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
