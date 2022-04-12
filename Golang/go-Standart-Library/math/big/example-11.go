package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	r := new(big.Rat)
	_, err := fmt.Sscan("1.5000", r)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(r)
	}
}
