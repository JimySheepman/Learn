package main

import "fmt"

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

type Doubly struct {
	Head *Node
}

func New() *Doubly {
	return &Doubly{}
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}

func (ll *Doubly) AddAtBeg(val interface{}) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = n
	n.Prev = cur
}

func (ll *Doubly) AddAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next

	if ll.Head != nil {
		ll.Head.Prev = nil
	}
	return cur.Val
}

func (ll *Doubly) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next

	if ll.Head != nil {
		ll.Head.Prev = nil
	}
	return cur.Val
}

func (ll *Doubly) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	return retval
}

func (ll *Doubly) Count() interface{} {
	var ctr int = 0

	for cur := ll.Head; cur != nil; cur = cur.Next {
		ctr += 1
	}
	return ctr
}

func (ll *Doubly) Reverse() {
	var Prev, Next *Node
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = Prev
		cur.Prev = Next
		Prev = cur
		cur = Next
	}

	ll.Head = Prev
}

func (ll *Doubly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, "")
	}
	fmt.Print("\n")
}

func (ll *Doubly) DisplayReverse() {
	if ll.Head == nil {
		return
	}

	var cur *Node
	for cur = ll.Head; cur.Next != nil; cur = cur.Next {
	}

	for ; cur != nil; cur = cur.Prev {
		fmt.Print(cur.Val, "")
	}
	fmt.Print("\n")
}
