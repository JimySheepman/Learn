package main

import "fmt"

// Concrete implementation
//
// Concrete Implementations contain platform-specific code.
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
