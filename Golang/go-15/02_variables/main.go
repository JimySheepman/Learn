package main

import "fmt"

func shorthand() {
	a := 10
	b := "golang"
	c := 4.14
	d := true
	e := `Do you like my hat?`
	f := 'M'

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

	fmt.Println()

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}

func var_zero_value() {
	var a int
	var b string
	var c bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)

	fmt.Println()

}

func less_emphasis() {
	// declare variable
	var message string
	message = "hello world!"
	fmt.Println(message)

	// declare many at one variables
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
}
func main() {
	shorthand()
	var_zero_value()
	less_emphasis()
}
