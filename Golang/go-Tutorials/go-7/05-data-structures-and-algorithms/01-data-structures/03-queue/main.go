package main

import "fmt"

func main() {
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Dequeue()
	queue.Enqueue(5)
	queue.Display()
}

type Queue struct {
	Data []interface{}
}

func New() *Queue {
	return &Queue{
		Data: []interface{}{},
	}
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.Data) == 0
}

func (queue *Queue) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.Data[0]
}

func (queue *Queue) Enqueue(item interface{}) *Queue {
	queue.Data = append(queue.Data, item)
	return queue
}

func (queue *Queue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	item := queue.Data[0]
	queue.Data = queue.Data[1:]

	return item

}

func (queue *Queue) Display() {
	for _, v := range queue.Data {
		fmt.Println(v)
	}
}
