package main

import (
	"encoding/xml"
	"fmt"
)

type MyXmlElement struct {
	XMLName xml.Name `xml:"myOwnName"`
	Name    string   `xml:"name"`
}

func main() {
	elem := MyXmlElement{Name: "Testing"}
	m, err := xml.MarshalIndent(elem, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(m))

}
