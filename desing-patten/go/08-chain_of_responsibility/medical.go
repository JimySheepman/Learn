package main

import "fmt"

// Concrete handler
type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicalDone {
		fmt.Println("Medical already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicalDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}
