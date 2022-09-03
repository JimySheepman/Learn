package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	myvar := os.Getenv("MYVAR")
	myvar2 := os.Getenv("MYVAR2")
	fmt.Printf("myvar : '%s'\n", myvar)
	fmt.Printf("myvar2 :'%s'\n", myvar2)

	port, found := os.LookupEnv("DB_PORT")
	if !found {
		log.Fatal("impossible to start up, DB_PORT env var is mandatory")
	}
	portParsed, err := strconv.ParseUint(port, 10, 8)
	if err != nil {
		log.Fatalf("impossible to parse db port: %s", err)
	}
	log.Println(portParsed)
}
