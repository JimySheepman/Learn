package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

func main() {
	var network bytes.Buffer

	gob.Register(Point{})

	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}

}

func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

func interfaceDecode(dec *gob.Decoder) Pythagoras {
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return p
}
