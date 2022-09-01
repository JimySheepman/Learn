package main

import (
	"encoding/xml"
	"fmt"
)

type MyXML struct {
	Cat `xml:"cat"`
}

type Cat struct {
	Name string `xml:"name"`
	Age  uint   `xml:"age"`
}

func main() {
	myXML := []byte(`<cat>
    <name>Ti</name>
    <age>23</age>
</cat>`)
	c := MyXML{}
	err := xml.Unmarshal(myXML, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)
}
