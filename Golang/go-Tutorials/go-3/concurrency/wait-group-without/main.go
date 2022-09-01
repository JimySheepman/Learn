package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Program start \n")
	for i := 0; i < 10; i++ {
		go concurrentTaks(i)
	}
	finishTask()
	fmt.Printf("Program end \n")

}

func finishTask() {
	fmt.Println("Executing finish task")
}

func concurrentTaks(taskNumber int) {
	fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d \n", taskNumber)
}
