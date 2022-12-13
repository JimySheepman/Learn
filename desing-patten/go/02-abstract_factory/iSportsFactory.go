package main

import "fmt"

// Abstract Factory interface
//
// The Abstract Factory interface declares a set of methods for creating each of the abstract products.
type ISportsFactory interface {
	makeShoe() IShirt
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
