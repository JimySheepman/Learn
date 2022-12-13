package main

// Product interface
//
// The Product declares the interface, which is common to all objects
// that can be produced by the creator and its subclasses.
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
