package main

import "fmt"

type Geometry interface {
	Edges() int
}

type Pentagon struct{}
type Hexagon struct{}
type Octagon struct{}
type Decagon struct{}

func (p Pentagon) Edges() int { return 5 }
func (h Hexagon) Edges() int  { return 6 }
func (o Octagon) Edges() int  { return 8 }
func (d Decagon) Edges() int  { return 10 }

func Parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}

func main() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	g := [...]Geometry{p, h, o, d}
	for _, v := range g {
		fmt.Println(Parameter(v, 5))
	}
}
