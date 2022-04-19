package main

import "fmt"

func main() {
	var action int
	fmt.Println("Enter 1 for student and 2 for professional")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Printf("I am a student")
	case 2:
		fmt.Println("I am a professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}
	fmt.Println("")
	fmt.Println("Enter 1 for US and 2 for UK")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("US")
	case 2:
		fmt.Println("UK")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}
}
