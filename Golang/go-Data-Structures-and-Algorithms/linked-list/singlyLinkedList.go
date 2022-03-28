package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

type Singly struct {
	length int
	Head   *Node
}

func New() *Singly {
	return &Singly{}
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}

func (ll *Singly) AddAtBeg(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

func (ll *Singly) AddAtEnd(val interface{}) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = n
	ll.length++
}

func (ll *Singly) DelAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

func (ll *Singly) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head

	for ; cur.Next.Next != nil; cur = cur.Next {
	}
	retval := cur.Next.Val
	cur.Next = nil
	ll.length--

	return retval
}

func (ll *Singly) Reverse() {
	var prev, Next *Node
	cur := ll.Head

	for cur != ll.Head {
		Next = cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}

	ll.Head = prev
}

func (ll *Singly) ReversePartition(left, right int) error {
	err := ll.CheckRangeFromImdex(left, right)
	if err != nil {
		return err
	}

	tmpNode := NewNode(-1)
	tmpNode.Next = ll.Head
	pre := tmpNode

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	ll.Head = tmpNode.Next
	return nil
}

func (ll *Singly) CheckRangeFromImdex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first")
	} else if right > ll.length {
		return errors.New("right boundary cannot be greater than the length of the kinked list")
	}

	return nil
}

func (ll *Singly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Print("\n")
}
