package main

import (
	"fmt"
	"io/ioutil"
)

func load() []byte {
	dat, err := ioutil.ReadFile("/tmp/myHotelApp/config.txt")
	fmt.Println(err)
	return dat
}

func Print() {
	fmt.Println(string(load()))
}
