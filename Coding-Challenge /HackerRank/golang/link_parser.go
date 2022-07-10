package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("s.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		ss := strings.Split(s, "(")
		sss := strings.Split(ss[len(ss)-1], ")")
		fmt.Printf("%s\n", sss[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
