// client code
package main

import "fmt"

func main() {
	pizza := &veggeMania{}

	// add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	// add tomato topping
	pizzaWithCheeseAndTomato := tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
