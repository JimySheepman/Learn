package main

import "fmt"

func init_condition_post() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func nested() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}

func for_condition_while() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func for_no_condition() {
	i := 0

	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 50 {
			break
		}
	}
}

func for_range() {
	for index, value := range []int{1, 23, 4, 5, 6, 1, 123} {
		fmt.Println("index: ", index, " value: ", value)
	}

	for _, value := range []int{1, 23, 4, 5, 6, 1, 123} {
		fmt.Println("value: ", value)
	}

	for index := range []int{1, 23, 4, 5, 6, 1, 123} {
		fmt.Println("index: ", index)
	}
}

func rune_loop() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}

func main() {
	init_condition_post()
	nested()
	for_condition_while()
	for_no_condition()
	for_range()
	rune_loop()
}
