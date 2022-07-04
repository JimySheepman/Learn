package main

import "fmt"

func main() {

	pointer()
	var a = 3
	doubleNotPointer(3)
	fmt.Println(a)

	doublePointer(&a)
	fmt.Println(a)
	p := &a
	doublePointer(p)
	fmt.Println(a, p == nil)
}

func pointer() {
	p0 := new(int)
	fmt.Println(p0)
	fmt.Println(*p0)

	x := *p0

	p1, p2 := &x, &x
	fmt.Println(p1 == p2)
	fmt.Println(p0 == p1)
	p3 := &*p0

	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3)

	fmt.Printf("%T, %T \n", *p0, x)
	fmt.Printf("%T, %T \n", p0, p1)
}

func doubleNotPointer(x int) {
	x += 3
}

func doublePointer(x *int) {
	*x += *x
	x = nil
}
