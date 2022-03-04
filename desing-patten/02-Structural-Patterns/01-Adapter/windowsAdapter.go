// Adapter
package main

import "fmt"

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}
