package main

import (
	"log"
	"strings"
)

func main() {
	EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
	upper(EUcountries)
	log.Println(EUcountries)
}

func upper(countries []string) {
	for k, _ := range countries {
		countries[k] = strings.ToUpper(countries[k])
	}
}
