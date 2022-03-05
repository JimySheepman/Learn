// Handler interface
package main

type department interface {
	execute(*patient)
	setNext(department)
}
