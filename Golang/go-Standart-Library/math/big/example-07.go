package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	f := new(big.Float)
	_, err := fmt.Sscan("1.19282e99", f)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(f)
	}
}
