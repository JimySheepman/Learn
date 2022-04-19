package main

import "fmt"

func if_statement() {
	// eval true
	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}

	// not exclamation
	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}

	// init statement
	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}

func if_else_statement() {
	if false {
		fmt.Println("first print statement")
	} else {
		fmt.Println("second print statement")
	}
}

func if_elseif_else_statement() {
	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("ahahaha print statement")
	} else {
		fmt.Println("third print statement")
	}
}

func switch_statement() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}

func multiple_switch_statement() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}
func no_expression_switch_statement() {

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}
func on_type_switch_statement() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

func main() {
	// if
	if_statement()
	if_else_statement()
	if_elseif_else_statement()

	// switch
	switch_statement()
	multiple_switch_statement()
	no_expression_switch_statement()
	on_type_switch_statement()
}
