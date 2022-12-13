package main

import "fmt"

// Refined abstraction
//
// Refined Abstractions provide variants of control logic.
// Like their parent, they work with different implementations
// via the general implementation interface.
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
