// arraylerin esnek hali -> slice
// kesilmiş bir arraay (aralıklı py'deki gibi -> slice
// map dictinory ile aynı şeydir (key - value) sistemlerdir.
package main

import "fmt"

func main() {
	myArray := [...]int{4, 5, 6, 7, 8, 1234, 12, 47, 4, 345, 976}
	mySlice := myArray[:]
	fmt.Println(mySlice)
	mySlice[0] = 11
	fmt.Println("array: ", myArray)
	fmt.Println("slice: ", myArray)

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 12
	numbers[2] = 123
	numbers[3] = 1234
	numbers[4] = 12345
	fmt.Println(numbers)
	numbers = append(numbers, 123456)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "foo"
	myMap[12] = "bar"
	fmt.Println(myMap)

	states := make(map[string]string)
	states["ISt"] = "istanbul"
	states["ANK"] = "ankara"
	states["ANT"] = "antalya"
	fmt.Println(states)

	antalya := states["ANT"]
	fmt.Println("Seçili şehir, ", antalya)

	delete(states, "ANK")
	fmt.Println(states)

	states["ERZ"] = "erzurum"

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

}
