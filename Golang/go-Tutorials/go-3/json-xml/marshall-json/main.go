package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID       uint64   `json:"id"`
	Name     string   `json:"name"`
	SKU      string   `json:"sku"`
	Category Category `json:"category"`
}
type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func main() {
	p := Product{
		ID:       42,
		Name:     "Tea Pot",
		SKU:      "TP12",
		Category: Category{ID: 2, Name: "Tea"}}
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	bI, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bI))
}
