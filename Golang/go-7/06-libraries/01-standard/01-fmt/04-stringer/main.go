package main

import (
	"fmt"
)

type Item struct {
	Id   int
	Name string
}

func (i Item) String() string {
	return fmt.Sprintf("[#%d] %s", i.Id, i.Name)
}

func (i Item) GoString() string {
	return fmt.Sprintf("{ id:%d, name:%q }", i.Id, i.Name)
}

func main() {
	Sample("Stringer", Stringer)
	Sample("GoStringer", GoStringer)
}

func Stringer() {
	item := Item{1, "Item 1"}

	fmt.Printf("String(): %v\n", item)
}

func GoStringer() {
	item := Item{2, "Item 2"}

	fmt.Printf("GoString(): %#v\n", item)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
