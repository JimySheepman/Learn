package main

import "fmt"

const p = "death & taxes"

func constantVariable() {
	const q = 42

	fmt.Println(p)
	fmt.Println(q)
}

const (
	pi       = 3.14
	language = "Go"
)

func multiple_initialization() {
	fmt.Println(pi)
	fmt.Println(language)
}

const (
	a = iota
	b = iota
	c = iota
	d
	e
	f = iota * 10
	g = iota * 10
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func defineIota() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println()

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}

func main() {
	constantVariable()
	multiple_initialization()
	defineIota()
}
