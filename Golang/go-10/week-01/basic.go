package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// Cosntant variable
const (
	Pi = 3.14
)

// Gloab Variable
var (
	m int
	n int
)

func main() {
	DataType()
	Operators()
	ConditionalStatements()
	Loops()
	Functions()
	Array()
	Slices()
	Maps()
	Struct()
	Interface()

}

func DataType() {
	// Data Type initialization and Declaration
	var x int = 1
	var y int
	fmt.Println(x, y)

	var a, b, c = 5.25, 25.25, 14.15
	fmt.Println(a, b, c)

	city := "Berlin"
	country := "Germany"
	fmt.Println(city, country)

	food, drink, price := "Pizza", "Pepsi", 125
	fmt.Println(food, drink, price)

	m, n = 1, 2
	fmt.Println(m, n)

	// Integer Float String Boolean
	// uint (0 to 255)
	var n1 uint8 = 255
	var n2 uint16 = 65535
	var n3 uint32 = 4294967295
	var n4 uint64 = 18446744073709551615

	fmt.Println(n1, n2, n3, n4)

	// int (-128 to 127)
	var n5 int = 255
	var n6 int16 = 65535 / 2
	var n7 int32 = 4294967295 / 2
	var n8 int64 = 18446744073709551615 / 2

	fmt.Println(n5, n6, n7, n8)

	// float & coplex IEEE-754
	var n9 int = 255
	var n10 int16 = 65535 / 2
	var n11 int32 = 4294967295 / 2
	var n12 int64 = 18446744073709551615 / 2

	fmt.Println(n9, n10, n11, n12)

	// boolean & string
	var boolean bool = true
	var str string = "string"

	fmt.Println(boolean, str)
}

func Operators() {
	// Arithmetic Operators
	var m, n int = 100, 25

	fmt.Println("m + n = ", m+n)
	fmt.Println("m - n = ", m-n)
	fmt.Println("m * n = ", m*n)
	fmt.Println("m / n = ", m/n)
	fmt.Println("m % n = ", m%n)

	// Bitwise Operators
	x := 75
	y := 25

	fmt.Println("bitwise AND -> x & y = ", x&y)
	fmt.Println("bitwise OR -> x | y = ", x|y)
	fmt.Println("bitwise XOR -> x ^ y = ", x^y)
	fmt.Println("bit clear AND NOT -> x &^ y = ", x&^y)

	// Comparison operators
	i, j := 60, 50

	fmt.Println("equal -> i == j = ", i == j)
	fmt.Println("not equal -> i != j = ", i != j)
	fmt.Println("less -> i < j = ", i < j)
	fmt.Println("less or equal -> i <= j = ", i <= j)
	fmt.Println("greater -> i > j = ", i > j)
	fmt.Println("greater or equal -> i >= j = ", i >= j)

	// Logical Operators
	if i != j && i >= j {
		fmt.Println("locial AND")
	}

	if i != j || i <= j {
		fmt.Println("locial OR")
	}

	// Assignment Operators
	a, b := 50, 60
	a += b
	fmt.Println(a)

	a -= b
	fmt.Println(a)

	a *= b
	fmt.Println(a)

	a /= b
	fmt.Println(a)

	a %= b
	fmt.Println(a)

	a++
	fmt.Println(a)

	a--
	fmt.Println(a)
}

func ConditionalStatements() {
	// if...else if...else Statement
	x := 100

	if x == 100 {
		fmt.Println(x)
	} else if x > 100 {
		fmt.Println("lower to 100")
	} else {
		fmt.Println("upper to 100")
	}

	// if statement initialization
	if y := 100; y == 100 {
		fmt.Println(y)
	}

	//switch Statement
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Clean your house.")
	case 15:
		fmt.Println("Today is 15th. Clean your house.")
	default:
		fmt.Println("Go to Party.")
	}

	// switch multiple cases Statement
	switch today.Day() {
	case 5, 10, 15:
		fmt.Println("Today is 10th. Clean your house.")
	default:
		fmt.Println("Go to Party.")
	}

	// switch fallthrough case Statement
	switch today.Day() {
	case 20:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

	// swith conditional cases Statement
	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}

	// switch initializer Statement
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func Loops() {
	// traditional for Statement
	for k := 1; k <= 10; k++ {
		fmt.Println(k)
	}

	// while for
	k := 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	// Infinite for
	for i := 0; ; i++ {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}

	// range for
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa", "Turkey": "Istanbul"}
	for index, value := range strDict {
		fmt.Println("index: ", index, " value: ", value)
	}

	// example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	// range loop over string
	for range "Hello" {
		fmt.Println("Hello")
	}

	// Infinite loop
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}

func Functions() {
	// Function with return value
	fmt.Println(addWithReturn(10, 5))

	// Function without return value
	addWithoutReturn(7, 12)

	// Return types can have names
	fmt.Println(addWithReturnHaveNames(19, 12))

	// Variadic Functions
	variadicExample("red", "blue", "green")
	variadicExample("red", "blue", "green", "yellow")

	// Normal function parameter with variadic function parameter
	fmt.Println(calculation("Square", 20))
	fmt.Println(calculation("Rectangle", 20, 30))

	// Pass different types of arguments in variadic function
	variadicExampleInterface(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}

// --- Example Funtions Start ---
func addWithReturn(x int, y int) int {
	return x + y
}

func addWithoutReturn(x, y int) {
	fmt.Println(x + y)
}

func addWithReturnHaveNames(x, y int) (result int) {
	result = x + y
	return
}

func variadicExample(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func calculation(str string, y ...int) int {
	area := 1
	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

func variadicExampleInterface(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

// --- Example Funtions Finish ---

func Array() {
	// A basic declaration of an array goes like this
	var intArray [5]int
	intArray[0] = 10
	fmt.Println(intArray)

	//Declare an array with some or all values to it at the time of declaration
	intArray = [5]int{10, 20, 30}
	fmt.Println(intArray)

	intArray = [5]int{0: 10, 2: 30, 4: 50}
	fmt.Println(intArray)

	// Declare an array inside function using the := shortcut
	intArray = [5]int{10, 20, 30, 40, 50}
	fmt.Println(intArray)

	// Declaring an array with ellipses ...
	intArray = [...]int{10, 20, 30, 40, 50}
	fmt.Println(intArray)

	// for loop to iterate over the Array
	intArray = [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	// Assigning an array by value and reference
	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	fmt.Printf("strArray1: %v\n", strArray1)
	strArray2 := strArray1 // data is passed by value (copied)
	fmt.Printf("strArray2: %v\n", strArray2)
	strArray1[0] = "Canada"
	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)

	strArray3 := &strArray1
	fmt.Printf("strArray3: %v\n", strArray3)
	fmt.Printf("&strArray3: %v\n", &strArray3)
	fmt.Printf("*strArray3: %v\n", *strArray3)

	// An example program that explores the array type
	var x [5]int // Array Declaration
	x[0] = 10    // Assign the values to specific Index
	x[4] = 20    // Assign Value to array index in any Order
	x[1] = 30
	x[3] = 40
	x[2] = 50
	fmt.Println("Values of Array X: ", x)

	// Array Declartion and Intialization to specific Index
	y := [5]int{0: 100, 1: 200, 3: 500}
	fmt.Println("Values of Array Y: ", y)

	// Array Declartion and Intialization
	Country := [5]string{"US", "UK", "Australia", "Russia", "Brazil"}
	fmt.Println("Values of Array Country: ", Country)

	// Array Declartion without length and Intialization
	Transport := [...]string{"Train", "Bus", "Plane", "Car", "Bike"}
	fmt.Println("Values of Array Transport: ", Transport)
}

func Slices() {
	// A basic declaration of an slice using make() function
	var intSlice1 = make([]int, 10)     // when length and capacity is same
	var intSlice2 = make([]int, 10, 20) // when length and capacity is different

	fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
	fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))

	// Declaration using new keyword
	intSlice1 = new([50]int)[0:10]

	fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))

	//  Literal declaration
	var intSlice = []int{10, 20, 30, 40}
	var strSlice1 = []string{"India", "Canada", "Japan"}

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice1), cap(strSlice1))
	fmt.Println(strSlice1)

	//Enlarge a slice using the append function
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20

	fmt.Println("Slice A:", a)
	fmt.Println("Slice len:", len(a), " Slice cap: ", cap(a))

	a = append(a, 30, 24, 40, 50, 1234, 41, 2)
	fmt.Println("Slice A:", a)
	fmt.Println("Slice len:", len(a), " Slice cap: ", cap(a))

	// Enlarge a slice using the copy function
	a = []int{5, 6, 7}
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

	b := make([]int, 5, 10)
	copy(b, a) // copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)

	//  Slice tricks
	var countries = []string{"india", "japan", "canada", "australia", "russia"}

	fmt.Printf("Countries: %v\n", countries)
	fmt.Printf(":2 %v\n", countries[:2])
	fmt.Printf("1:3 %v\n", countries[1:3])
	fmt.Printf("2: %v\n", countries[2:])
	fmt.Printf("2:5 %v\n", countries[2:5])
	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[4])
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last element: %v\n", countries[4:])
	fmt.Printf("All elements: %v\n", countries[0:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])

	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	// Assign parts of slice to another slice
	var oldStr = []string{"india", "japan", "canada", "australia", "russia"}
	var strSlice []string
	newStr := oldStr[0:3]

	strSlice = append(strSlice, oldStr[:1]...)
	fmt.Printf("newStr: %v\n", newStr)
	fmt.Printf("strSlice: %v\n", strSlice)

	fmt.Printf("oldStr length: %v\tcapacity: %v\n", len(oldStr), cap(oldStr))
	fmt.Printf("newStr length: %v\tcapacity: %v\n", len(newStr), cap(newStr))
	fmt.Printf("strSlice length: %v\tcapacity: %v\n", len(strSlice), cap(strSlice))

	newStr[0] = "china"
	fmt.Printf("newStr: %v\n", newStr)
	fmt.Printf("oldStr: %v\n", oldStr)

	newStr = append(newStr, "brazil")
	fmt.Printf("newStr: %v\n", newStr)
	fmt.Printf("oldStr: %v\n", oldStr)

	oldStr = append(oldStr, "us")
	newStr = append(newStr, "uk")
	fmt.Printf("newStr: %v\n", newStr)
	fmt.Printf("oldStr: %v\n", oldStr)

	// Append a slice to an existing slice
	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}
	slice2 = append(slice2, slice1...)

	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	// Append part of slice to an existing slice
	slice1 = []string{"india", "japan", "canada", "us", "uk", "italy", "germany"}
	slice2 = []string{"australia", "russia"}
	slice2 = append(slice2, slice1[3:]...)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
}

func Maps() {
	// Empty Map declaration
	var employee = map[string]int{}

	fmt.Println(employee)

	// Map initialization
	employee = map[string]int{"Mark": 10, "Sandy": 20}

	fmt.Println(employee)

	// Creating Map using the make function
	employee = make(map[string]int)
	employee["Mark"] = 20
	employee["Sandy"] = 30
	fmt.Println(employee)

	// Find length of Map
	employeeList := make(map[string]int)

	fmt.Println(len(employee))
	fmt.Println(len(employeeList))

	// Delete element from a Map
	delete(employee, "Mark")
	fmt.Println(employee)

	// Adding and Editing elements in Map
	employee["Rocky"] = 30 // Add element
	employee["Josef"] = 40
	employee["Mark"] = 50 // Edit element
	fmt.Println(employee)

	// for range loop to iterate over a Map
	for key, value := range employee {
		fmt.Println("Key: ", key, "  Value: ", value)
	}
}

// --- Example Struct Start ---
type rectangle struct {
	length  float64
	breadth float64
	color   string
}

type rectangle2 struct {
	length   int
	breadth  int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
}

type Employee1 struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func Struct() {
	// Struct Declaration
	fmt.Println(rectangle{10.5, 25.10, "red"})

	// Dot notation
	var rect rectangle2
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)

	// var keyword and := operator
	rect1 := new(rectangle)
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"

	fmt.Println(rect1)

	var rect2 = new(rectangle)
	rect2.length = 10
	rect2.color = "Red"

	fmt.Println(rect2) // breadth skipped

	// Using &
	rect1 = &rectangle{10, 20, "Green"} // Can't skip any value

	fmt.Println(rect1)

	rect2 = &rectangle{}
	rect2.length = 10
	rect2.color = "Red"

	fmt.Println(rect2) // breadth skipped

	var rect3 = &rectangle{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"

	fmt.Println(rect3) // length skipped

	// Use field tags in the defination of Struct type
	json_string := `
	{
	  "firstname": "Rocky",
	  "lastname": "Sting",
	  "city": "London"
	}`

	emp1 := new(Employee1)
	json.Unmarshal([]byte(json_string), emp1)

	fmt.Println(emp1)

	emp2 := new(Employee1)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)

	fmt.Printf("%s\n", jsonStr)

	// Nested Struc Type
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])

	// Adding Methods to Struct Types
	e = Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}

// Method
func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

// --- Example Struct Finish ---

// --- Example Interface Start ---
type Information interface {
	General()
	Attributes()
	Inventory()
}

type Product struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func (prd Product) General() {
	fmt.Printf("\n%s", prd.Name)
	fmt.Printf("\n%s\n", prd.Description)
	fmt.Println(prd.Price)
}

func (prd Product) Attributes() {
	fmt.Println(prd.Weight)
}

type Mobile struct {
	Product
	DisplayFeatures   []string
	ProcessorFeatures []string
}

// Method Overriding
func (mob Mobile) Attributes() {
	mob.Product.Attributes()
	fmt.Println("\nDisplay Features:")
	for _, key := range mob.DisplayFeatures {
		fmt.Println(key)
	}
	fmt.Println("\nProcessor Features:")
	for _, key := range mob.ProcessorFeatures {
		fmt.Println(key)
	}
}

func Interface() {
	// mplemention of Interface into Concrete Types
}
