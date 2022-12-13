package main

import "fmt"

// Refined abstraction
//
// Refined Abstractions provide variants of control logic.
// Like their parent, they work with different implementations
// via the general implementation interface.
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
