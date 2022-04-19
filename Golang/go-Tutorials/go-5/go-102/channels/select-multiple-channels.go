package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 2)
	chan2 <- 2
	/*
		select {
			case c1Val := <-chan1:
				fmt.Println(c1Val)
			case c2Val := <-chan2:
				fmt.Println(c2Val)
		}

		for {
			select {
				case c1Val := <-chan1:
					fmt.Println(c1Val)
				case c2Val := <-chan2:
					fmt.Println(c2Val)
			}
		}

		for {
			select {
				case c1Val := <-chan1:
					fmt.Println(c1Val)
				case c2Val := <-chan2:
					fmt.Println(c2Val)
				default:
					break
			}
		}
	*/
	var done bool
	for !done {
		select {
		case c1Val := <-chan1:
			fmt.Println(c1Val)
		case c2Val := <-chan2:
			fmt.Println(c2Val)
		default:
			done = true
		}
	}

}
