package stack

type Node struct {
	Val  interface{}
	Next *Node
}

type Stack struct {
	top    *Node
	length int
}

func New() *Stack {
	return &Stack{}
}

func (ll *Stack) push(n interface{}) {
	newStack := &Node{}

	newStack.Val = n
	newStack.Next = ll.top

	ll.top = newStack
	ll.length++
}

func (ll *Stack) pop() interface{} {
	result := ll.top.Val
	if ll.top.Next == nil {
		ll.top = nil
	} else {
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}
	ll.length--
	return result
}

func (ll *Stack) isEmpty() bool {
	return ll.length == 0
}

func (ll *Stack) len() int {
	return ll.length
}
func (ll *Stack) show() (in []interface{}) {
	current := ll.top

	for current != nil {
		in = append(in, current.Val)
		current = current.Next
	}
	return
}
