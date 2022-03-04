// Client code
package main

import "fmt"

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client insert Lightning connector into computer.")
	com.insertIntoLightningPort()
}
