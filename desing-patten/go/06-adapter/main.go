package main

// The client code doesn’t get coupled to the concrete adapter class
// as long as it works with the adapter via the client interface.
// Thanks to this, you can introduce new types of adapters into
// the program without breaking the existing client code. This can be
// useful when the interface of the service class gets changed or replaced:
// you can just create a new adapter class without changing the client code.
func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
