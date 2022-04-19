package main

import (
	"fmt"
	"os"
)

func main() {
	Sample("Scan", Scan)
	Sample("Scanln", Scanln)
	Sample("Scanf", Scanf)
	Sample("Sscan", Sscan)
	Sample("Sscanln", Sscanln)
	Sample("Sscanf", Sscanf)
	Sample("Fscan", Fscan)
	Sample("Fscanln", Fscanln)
	Sample("Fscanf", Fscanf)
}

func Scan() {
	// read with space-separated
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Scan(&age)

	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Scanln() {
	// read with lines
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)

	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Scanf() {
	// read with format
	var name string
	var age int
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Sscan() {
	// reads string same as Scan

	var i1, i2 int
	fmt.Sscan("100 200", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Sscanln() {
	// reads string same as Scanln

	var i1, i2 int
	fmt.Sscanln("100\n200", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Sscanf() {
	// reads string same as Scanf

	var i1, i2 int
	fmt.Sscanf("100 and 200", "%d and %d", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Fscan() {
	// reads from reader same as Scan

	var i1, i2 int
	fmt.Fscan(os.Stdin, &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Fscanln() {
	// reads from reader same as Scanln

	var i1, i2 int
	fmt.Fscanln(os.Stdin, &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Fscanf() {
	// reads from reader same as Scanf

	var i1, i2 int
	fmt.Fscanf(os.Stdin, "%d and %d", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
