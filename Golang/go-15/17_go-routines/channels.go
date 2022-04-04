package main

import (
	"fmt"
	"time"
)

func unbuffered_channels_block() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(1 * time.Second)
}

func channels_range() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}

func n_to_one() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

func one_to_n() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}

func pass_return_channels() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func closure_bindin() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	for _ = range values {
		<-done
	}
}

func deadlock_challenges() {
	/* problem
	   c := make(chan int)
	   c <- 1
	   fmt.Println(<-c)
	*/

	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

}

func factorial_challenge() {
	/* 	f := factorial(4)
	   	fmt.Println("Total:", f)
	*/

	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

/* func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
} */

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func main() {
	unbuffered_channels_block()
	fmt.Println("------------------------")
	channels_range()
	fmt.Println("------------------------")
	n_to_one()
	fmt.Println("------------------------")
	one_to_n()
	fmt.Println("------------------------")
	pass_return_channels()
	fmt.Println("------------------------")
	closure_bindin()
	fmt.Println("------------------------")
	deadlock_challenges()
	fmt.Println("------------------------")
	factorial_challenge()
}
