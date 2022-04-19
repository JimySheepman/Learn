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
		fmt.Printf("I am a professional")
	default:
		panic(fmt.Sprintf("I am a %d", action))
	}
	defer func() {
		action := recover()
		fmt.Println(action)
	}()
}
