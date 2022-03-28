package main

import "fmt"

const arraySize = 10

type Stack struct {
	top  int
	data [arraySize]int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(i int) bool {
	if s.top == len(s.data) {
		return false
	}

	s.data[s.top] = i
	s.top += 1
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	i := s.data[s.top-1]
	s.top -= 1
	return i, true
}

func (s *Stack) Peek() int {
	return s.data[s.top-1]
}

func (s *Stack) Get() []int {
	return s.data[:s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) Empty() {
	s.top = 0
	s.data = [arraySize]int{}
}

func main() {
	stack := New()
	stack.Push(1)
	fmt.Println(stack.data)
	stack.Push(4)
	stack.Push(5)
	stack.Push(2)
	stack.Push(7)
	fmt.Println(stack.data)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.data)
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
	stack.Empty()
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.data)
}
