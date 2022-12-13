package main

//	Abstraction
//
// The Abstraction provides high-level control logic.
// It relies on the implementation object to do the actual low-level work.
type Computer interface {
	Print()
	SetPrinter(Printer)
}
