package main

// Handler interface
type Department interface {
	execute(*Patient)
	setNext(Department)
}
