package main

import "fmt"

// Client
//
// The Client is a class that contains the existing business logic of the program.
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
