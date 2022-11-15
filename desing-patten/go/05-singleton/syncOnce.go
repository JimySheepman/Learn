package main

import (
	"fmt"
	"sync"
)

var once sync.Once

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
