package main

import "fmt"

func init_map() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)

	myGreetingMake := make(map[string]string)
	myGreetingMake["Tim"] = "Good"
	fmt.Println(myGreetingMake)

	myGreetingShort := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreetingShort["Ahmet"] = "Günaydın"
	fmt.Println(myGreetingShort)
	fmt.Println(len(myGreetingShort))
	myGreetingShort["Ahmet"] = "Gunaydin"
	fmt.Println(myGreetingShort)
	delete(myGreetingShort, "Ahmet")
	fmt.Println(myGreetingShort)
	fmt.Println()
}

func comma_ok_idiom() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	// delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	for key, val := range myGreeting {
		fmt.Println("key: ", key, "\tvalue: ", val)
	}

}

func main() {
	init_map()
	comma_ok_idiom()
}
