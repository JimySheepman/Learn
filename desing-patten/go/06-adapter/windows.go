package main

import "fmt"

// Unknown Service
//
// The Service is some useful class (usually 3rd-party or legacy).
// The client can’t use this class directly because it has an incompatible interface.
type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
