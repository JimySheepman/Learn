package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	time.Sleep(1 * time.Second)
	fmt.Println("---")
	wg.Add(2)
	go SayGreetingsWaitGroup("hi!", 10)
	go SayGreetingsWaitGroup("hello!", 10)
	wg.Wait()

	deferFuntion()
	recoverFunction()
	panicFunction()
}

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
}

func SayGreetingsWaitGroup(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done()
}

func deferFuntion() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")
}

func goFunctionValue() {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a)
	}(a)
	a = 769
	time.Sleep(time.Second * 2)
}

func recoverFunction() {
	defer func() {
		fmt.Println("exit normally")
	}()
	fmt.Println("hi!")

	defer func() {
		v := recover()
		fmt.Println("recovered: ", v)
	}()
	panic("bye!")
	fmt.Println("unreachable")
}

func panicFunction() {
	fmt.Println("hi!")
	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()
	for {
		time.Sleep(time.Second)
	}
}
