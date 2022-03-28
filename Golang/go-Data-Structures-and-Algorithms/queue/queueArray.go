package queue

var ListQueue []interface{}

func EnQueue(n interface{}) {
	ListQueue = append(ListQueue, n)
}

func DeQueue() interface{} {
	data := ListQueue[0]
	ListQueue = ListQueue[1:]
	return data
}

func FrontQueue() interface{} {
	return ListQueue[0]
}

func BackQueue() interface{} {
	return ListQueue[len(ListQueue)-1]
}

func LenQueue() int {
	return len(ListQueue)
}

func IsEmptyQueue() bool {
	return len(ListQueue) == 0
}
