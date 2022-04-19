package main

import "fmt"

func main() {
	var num1 int = 5
	fmt.Println(num1)
	num2 := 3
	fmt.Println(num2)
	num1 += num2
	fmt.Println(num1)
	num1 -= num2
	fmt.Println(num1)
	num1 *= num2
	fmt.Println(num1)
	num1 /= num2
	fmt.Println(num1)
	num1 %= num2
	fmt.Println(num1)
	fmt.Println("Relational operators")
	var result bool = false
	result = 1 == 1
	fmt.Println("1==1", result)
	result = 1 != 2
	fmt.Println("1!=2", result)
	result = 1 > 2
	fmt.Println("1>2", result)
	result = 1 < 2
	fmt.Println("1<2", result)
	result = 1 >= 1
	fmt.Println("1>=2", result)
	result = 1 <= 1
	fmt.Println("1<=2", result)
	fmt.Println("Logical operators")
	var op1, op2 bool = true, false
	fmt.Println("op1 AND op2", op1 && op2)
	fmt.Println("op1 OR op2", op1 || op2)
	fmt.Println("op1 NOT true", !op1)
	fmt.Println("op2 NOT true", !op2)
}
