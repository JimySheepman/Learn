package main

// Client interface
//
// The Client Interface describes a protocol that other classes must
// follow to be able to collaborate with the client code.
type Computer interface {
	InsertIntoLightningPort()
}
