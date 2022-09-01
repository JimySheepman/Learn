package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Program start \n")
	// initialize the wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go concurrentTaks(i, &waitGroup)
	}
	waitGroup.Wait()
	finishTask()
	fmt.Printf("Program end \n")

}

func finishTask() {
	fmt.Println("Executing finish task")
}

func concurrentTaks(taskNumber int, waitGroup *sync.WaitGroup) {
	fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d \n", taskNumber)
	waitGroup.Done()
}
