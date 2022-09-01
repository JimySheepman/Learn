package main

import (
	"encoding/xml"
	"fmt"
)

type Product struct {
	Name string `xml:"first>second>third"`
}

func main() {
	c := Product{}
	c.Name = "testing"
	b, err := xml.MarshalIndent(c, "", "        ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
