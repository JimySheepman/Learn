package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700
	printBook(Book1)
	printBook(Book2)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println("numbers ==", numbers)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	fmt.Println("numbers[:3] ==", numbers[:3])
	fmt.Println("numbers[4:] ==", numbers[4:])
	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	number2 := numbers[:2]
	printSlice(number2)
	number3 := numbers[2:5]
	printSlice(number3)

	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
