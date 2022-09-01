package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("myFile.test")
	if err != nil {
		fmt.Println(err)
		return
	}

	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info.Mode())

	fmt.Printf("Mode Numeric : %o\n", info.Mode())
	fmt.Printf("Mode Symbolic : %s\n", info.Mode())

	err = f.Chmod(0777)
	if err != nil {
		fmt.Println(err)
	}
}
