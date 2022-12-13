package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// Singleton
//
// The Singleton class declares the static method getInstance that returns
// the same instance of its own class. The Singleton’s constructor should be
// hidden from the client code. Calling the getInstance method should be
// the only way of getting the Singleton object.
type singleOnce struct {
}

var singleOnceInstance *singleOnce

// Singleton
func getOnceInstance() *singleOnce {
	if singleOnceInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleOnceInstance = &singleOnce{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleOnceInstance
}
