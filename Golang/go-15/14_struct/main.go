package main

import "fmt"

// * alias type defined types
type foo int

// * struct fields vvalues initialization
type person struct {
	first string
	last  string
	age   int
}

// * methods
func (p person) fullName() string {
	return p.first + p.last
}

// * embedded types
type doubleZero struct {
	person
	LicenseToKill bool
}

// * overriding fields & methods

type persons struct {
	First string
	Last  string
	Age   int
}

type doubleZeros struct {
	persons
	First         string
	LicenseToKill bool
}

func (p persons) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZeros) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

// * struct pointer
type personPointer struct {
	name string
	age  int
}

func main() {
	// * alias type defined types
	var myAge foo
	myAge = 44
	fmt.Printf("%T \t%v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T \t%v \n", yourAge, yourAge)

	// * struct fields vvalues initialization
	p1 := person{"James ", "Bond", 20}
	p2 := person{"Miss ", "Moneypenny", 19}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	// * methods
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	// * embedded types
	p3 := doubleZero{
		person: person{
			first: "James ",
			last:  "Bound",
			age:   21,
		},
		LicenseToKill: true,
	}
	fmt.Println(p3.first, p3.last, p3.age, p3.LicenseToKill)

	// * overriding fields
	p4 := doubleZeros{
		persons: persons{
			First: "James ",
			Last:  "Bound",
			Age:   12,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}
	fmt.Println(p4.First, p4.persons.First)
	p4.Greeting()
	p4.persons.Greeting()

	// * struct pointer
	p5 := &personPointer{"James", 20}
	fmt.Println(p5)
	fmt.Printf("%T \n", p5)
	fmt.Println(p5.name)
	fmt.Println(p5.age)
}
