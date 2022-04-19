package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())

	projects := make(chan string, 10)

	goRoutine.Add(5)
	for i := 1; i <= 5; i++ {
		go employee(projects, i)
	}
	for j := 1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project: %d", j)
	}

	close(projects)
	goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()

	for {
		project, result := <-projects

		if result == false {
			fmt.Printf("Employee: %d : Exit\n", employee)
			return
		}

		fmt.Printf("Employee: %d : Started %s\n", employee, project)

		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("\nTime to sleep", sleep, "ms\n")

		fmt.Printf("Employee: %d : Complated %s\n", employee, project)
	}
}
