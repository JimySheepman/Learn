package main

import "fmt"

// Concrete implementation
//
// Concrete Implementations contain platform-specific code.
type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
