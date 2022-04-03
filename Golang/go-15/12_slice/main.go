package main

import "fmt"

func init_slice() {
	var mySlice = []int{1, 3, 5, 7, 9}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	for i, v := range mySlice {
		fmt.Println(i, " - ", v)
	}
}

func calp_len_slice() {
	mySlice := make([]int, 1, 3)
	mySlice[0] = 1
	fmt.Println("My slice: ", mySlice, "\tSlice len: ", len(mySlice), "\tSlice cap: ", cap(mySlice))
	fmt.Println()

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "\tCapacity:", cap(mySlice), "\tValue: ", mySlice[i])
	}
}

func slicing_a_slice() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[3:])
	fmt.Println(mySlice[:])
	fmt.Println("myString"[2])
}

func append_to_slice() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)

}

func main() {
	init_slice()
	calp_len_slice()
	slicing_a_slice()
	append_to_slice()
}
