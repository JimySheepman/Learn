package main

import "fmt"

// Adapter
//
// The Adapter is a class that’s able to work with both the client and the service:
// it implements the client interface, while wrapping the service object.
// The adapter receives calls from the client via the adapter interface and
// translates them into calls to the wrapped service object in a format it can understand.
type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
