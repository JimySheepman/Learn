package main

//	Implementation
//
// The Implementation declares the interface that’s common for
// all concrete implementations. An abstraction can only communicate
// with an implementation object via methods that are declared here.
//
// The abstraction may list the same methods as the implementation,
// but usually the abstraction declares some complex behaviors that
// rely on a wide variety of primitive operations declared by the
// implementation.
type Printer interface {
	PrintFile()
}
