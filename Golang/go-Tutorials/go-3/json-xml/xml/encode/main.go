package main

import (
	"encoding/xml"
	"fmt"
)

type Product struct {
	ID       uint64   `xml:"id"`
	Name     string   `xml:"name"`
	SKU      string   `xml:"sku"`
	Price    float64  `xml:"price"`
	Category Category `xml:"category"`
}
type Category struct {
	ID   uint64 `xml:"id"`
	Name string `xml:"name"`
}

func main() {
	p := Product{ID: 42, Name: "Tea Pot", SKU: "TP12", Price: 30.5, Category: Category{ID: 2, Name: "Tea"}}
	bI, err := xml.MarshalIndent(p, "", "   ")
	if err != nil {
		panic(err)
	}
	xmlWithHeader := xml.Header + string(bI)
	fmt.Println(xmlWithHeader)
}
