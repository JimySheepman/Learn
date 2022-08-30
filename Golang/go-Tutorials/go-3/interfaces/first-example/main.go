package main

import "fmt"

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
	Stringer
}

type Stringer interface {
	String() string
}

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}

func (h Human) String() string {
	return fmt.Sprintf("human named %s %s of age %d living in %s", h.Firstname, h.Lastname, h.Age, h.Country)
}

type Cat struct {
	Name string
}

func (c Cat) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, from.Firstname)
}

func (c Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Firstname)
}
func (c Cat) String() string { return "" }

type Dog struct {
	Name string
}

func (d Dog) ReceiveAffection(from Human) {
	fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}

func (d Dog) String() string { return "" }

type Snake struct {
	Name string
}

func (s Snake) ReceiveAffection(from Human) {
	fmt.Printf("The snake named %s has received affection from Human named %s\n", s.Name, from.Firstname)
}

func (s Snake) GiveAffection(to Human) {
	fmt.Printf("The snake named %s has given affection to Human named %s\n", s.Name, to.Firstname)
}

func (s Snake) String() string { return "" }

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.ReceiveAffection(human)
}

func main() {
	var john Human
	john.Firstname = "John"
	john.Lastname = "Doe"
	john.Country = "USA"
	john.Age = 45

	fmt.Println(john)

	var c Cat
	c.Name = "Maru"

	var d Dog
	d.Name = "Medor"

	var snake Snake
	snake.Name = "Joe"

	Pet(c, john)
	Pet(d, john)
	Pet(snake, john)
}
