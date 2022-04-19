package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		consoleLog()
	}
	consoleLogParam("Jimmy")
	whoAmI("merlins", 31)
	result := add(5, 3)
	fmt.Println("\n", result)
	msg1, msg2 := swap("Hello World", "Hello There")
	fmt.Printf("msg1: %s\nmsg2: %s", msg1, msg2)
	helloWorld := func(msg string) {
		fmt.Println(msg)
	}
	helloWorld("\nHello, World!")

	func(msg string) {

		fmt.Println(msg)
	}("Hello")

	// function as value
	consoleLog2 := returnLambda()
	// invoke lambda that is returned
	consoleLog2("Hello")
}
func consoleLog() {
	fmt.Println("Hello from inside a function")
}
func consoleLogParam(msg string) {
	fmt.Println(msg)
}
func whoAmI(name string, age int) {
	fmt.Printf("Hi, your name is %s and you %d years odl.", name, age)
}
func add(num1 int, num2 int) int {
	return num1 + num2
}
func swap(a string, b string) (string, string) {

	return b, a
}

func returnLambda() func(string) {

	// return lambda
	return func(msg string) {

		fmt.Println(msg)
	}
}
