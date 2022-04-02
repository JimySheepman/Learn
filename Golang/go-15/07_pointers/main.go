package main

import "fmt"

// 				memory
// 				------
// 			 	\    \
// 		a = 43 	\ 43 \ -> 0x20818a220 <- &a
// 		*b = a	\    \ -> 0x20818a220 <- b = &a looking at the same address
// 				------
// 				\  . \
// 				\  . \
// 				\  . \
// 				\  . \
// 				\  . \
// 				------

func pointer() {
	// referencing
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	// dereferencing
	fmt.Println(*b)

	// using pointer
	*b = 42
	fmt.Println(a)
}

func zero_without_pointer(z int) {
	z = 0
}

func zero_with_pointer(z int) {
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	z = 0
}

func using_pointer() {
	// non pointer
	x := 5
	zero_without_pointer(x)
	fmt.Println(x)

	// with pointer
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	zero_with_pointer(x)
	fmt.Println(x)
}

func main() {
	pointer()
	using_pointer()
}
