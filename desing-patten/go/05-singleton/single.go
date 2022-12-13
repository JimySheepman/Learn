package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

// Singleton
//
// The Singleton class declares the static method getInstance that returns
// the same instance of its own class. The Singleton’s constructor should be
// hidden from the client code. Calling the getInstance method should be
// the only way of getting the Singleton object.
type single struct {
}

var singleInstance *single

// Singleton
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
