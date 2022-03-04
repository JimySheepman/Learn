// Unknown service
package main

import "fmt"

type windows struct {
}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB conncetor is plugged into windows machine.")
}
