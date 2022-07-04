package main

import (
	"fmt"
	"time"
)

func main() {
	// Ex1()
	// Ex2() // panic
	// Ex3() // infinity
	// Ex4() // fibonacci
	Ex5()

}

func Ex1() {
	c := make(chan int)
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		ch <- x * x
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)

	<-done
	fmt.Println("bye")
}

func Ex2() {
	c := make(chan int, 2)
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c))
	x, ok := <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
	x, ok = <-c
	fmt.Println(x, ok)
	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
	close(c)
	c <- 7
}

func Ex3() {
	var ball = make(chan string)
	kickBall := func(plaerName string) {
		for {
			fmt.Println(<-ball, "kicked the ball.")
			time.Sleep(time.Second)
			ball <- plaerName
		}
	}

	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	go kickBall("Emily")
	ball <- "referee"
	var c chan bool
	<-c
}

func Ex4() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y {
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}

	c := fibonacci()

	for x, ok := <-c; ok; x, ok = <-c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}

func Ex5() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default:
		}
	}
	tryRecive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-"
		}
	}

	trySend("Hello!")
	trySend("Hi!")
	trySend("Bye!")

	fmt.Println(tryRecive())
	fmt.Println(tryRecive())
	fmt.Println(tryRecive())
}
