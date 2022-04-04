package main

import "fmt"

// param argument
func greet(name string) {
	fmt.Println(name)
}

// multi param
func greet_multi_param(name, surname string) {
	fmt.Println(name, surname)
}

// return value
func greet_return(name, surname string) string {
	return fmt.Sprint(name, surname)
}

// return multiple value
func greet_multiple_return(name, surname string) (string, string) {
	return fmt.Sprint(name, surname), fmt.Sprint(name, surname)
}

// variadic params
func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// callback
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

// recursion
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// defer
func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

// pointer
func changeMe(z *int) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 24
}

func main() {
	// param argument
	greet("Jane")
	greet("John")

	// multi param
	greet_multi_param("John", "Doe")

	// return value
	fmt.Println(greet_return("John", "Doe"))

	// return multiple value
	fmt.Println(greet_multiple_return("John", "Doe"))

	// variadic params
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)

	// closure
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}

	// callback
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	// recursion
	fmt.Println(factorial(4))

	// defer
	defer world()
	hello()

	// anon self executing
	func() {
		fmt.Println("I'm driving!")
	}()

	// pointer
	age := 44
	fmt.Println(&age) // 0x82023c080
	changeMe(&age)
	fmt.Println(&age) //0x82023c080
	fmt.Println(age)  //24

}
