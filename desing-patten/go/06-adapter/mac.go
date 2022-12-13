package main

import "fmt"

// Service
//
// The Service is some useful class (usually 3rd-party or legacy).
// The client can’t use this class directly because it has an incompatible interface.
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
