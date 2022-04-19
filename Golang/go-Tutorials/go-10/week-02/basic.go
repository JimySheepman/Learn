package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"sync"
	"time"
)

func main() {
	DeferredFunctionsCalls()
	WithoutDefer()
	WithDefer()
	PanicandRecover()
	Concurrency()
	WorkingwithChannels()
	Logging()
}

func DeferredFunctionsCalls() {
	defer second()
	first()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func WithoutDefer() {
	file, err := os.Open("file")

	if err != nil {
		file.Close()
	}

	if file == nil {
		file.Close()
	}
	file.Close()
}

func WithDefer() {
	file, err := os.Open("file")
	defer file.Close()

	if err != nil {
		// do someting
	}

	if file == nil {
		// do someting
	}
}

func PanicandRecover() {
	var action int

	fmt.Println("Enter 1 for student and 2 for professional")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Printf("I am  a student")
	case 2:
		fmt.Printf("I am a professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}

	fmt.Println("")
	fmt.Println("Enter 1 for US and 2 for UK")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Printf("US")
	case 2:
		fmt.Printf("UK")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in:", err)
		}
	}()

	fmt.Println("Enter 1 for student and 2 for professional")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Printf("I am  a student")
	case 2:
		fmt.Printf("I am a professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}
}

type testConcurrency struct {
	min     int
	max     int
	country string
}

func printCountry(test *testConcurrency, groupTest *sync.WaitGroup) {
	for i := test.max; i > test.min; i-- {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(test.country)
	}

	fmt.Println("")
	groupTest.Done()
}

func Concurrency() {
	groupTest := new(sync.WaitGroup)

	japan := new(testConcurrency)
	china := new(testConcurrency)
	india := new(testConcurrency)

	japan.country = "Japan"
	japan.min = 0
	japan.max = 5

	china.country = "China"
	china.min = 0
	china.max = 5

	india.country = "India"
	india.min = 0
	india.max = 5

	go printCountry(japan, groupTest)
	go printCountry(china, groupTest)
	go printCountry(india, groupTest)

	groupTest.Add(3)
	groupTest.Wait()
}

var goRoutine sync.WaitGroup

func WorkingwithChannels() {
	unbuffered := make(chan int)
	buffered := make(chan int, 10)
	fmt.Println(unbuffered, buffered)

	rand.Seed(time.Now().Unix())

	projects := make(chan string, 10)

	goRoutine.Add(5)
	for i := 1; i <= 5; i++ {
		go employee(projects, i)
	}

	for j := 0; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :d%", j)
	}

	close(projects)
	goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		project, result := <-projects

		if result == false {
			fmt.Printf("Employee : %d : Exit\n", employee)
			return
		}

		fmt.Printf("Employee : %d : Started   %s\n", employee, project)

		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("\nTime to sleep", sleep, "ms")

		fmt.Printf("Employee : %d : Complated %s\n", employee, project)
	}
}

func Logging() {
	log.Println("mian started")

	fmt.Println("")
	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()

	log.Panicln("panic message")
	log.Fatalln("fatal message")
}

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
