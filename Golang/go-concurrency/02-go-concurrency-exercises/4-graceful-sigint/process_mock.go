package main

import (
	"fmt"
	"log"
	"time"
)

type MockProcess struct {
	isRunning bool
}

func (p *MockProcess) Run() {
	p.isRunning = true

	fmt.Print("Process running..")
	for {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
}

func (p *MockProcess) Stop() {
	if !p.isRunning {
		log.Fatal("Cannot stop a process which is not running")
	}

	fmt.Print("\nStopping process..")
	for {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
}
