package main

// Products
//
// Products are resulting objects. Products constructed by different
// builders don’t have to belong to the same class hierarchy or interface.
type House struct {
	windowType string
	doorType   string
	floor      int
}
