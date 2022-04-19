package main

import "fmt"

func main() {
	var groceryList = []string{

		"Bread",
		"Milk",
		"Cheese",
	}
	fmt.Println(groceryList[2])

	for i := 0; i < 3; i++ {
		fmt.Println(groceryList[i])
	}
	groceryList[0] = "Apples"
	for i, item := range groceryList {
		fmt.Printf("[%d] %s\n", i, item)
	}
	var employee = [2][2]string{

		{"John Doe", "Jane Bread"},
		{"Analyst", "Sysadmin"},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(employee[i][j])
		}
	}
	var greeting = []string{
		"Hello there",
		"Greetings",
	}
	ConsoleLog(greeting)
}

func ConsoleLog(msgs []string) {
	alen := len(msgs)
	for i := 0; i < alen; i++ {
		fmt.Println(msgs[i])
	}
}
