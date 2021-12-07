package main

import "fmt"

func main() {
	xSize := 16
	ySize := 16
	zSize := 128
	sum := 3
	for sum < 100 {
		fmt.Println(sum)
		sum += sum
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("i: ", i)
	}

	for x := 0; x <= xSize; x++ {
		for y := 0; y <= ySize; y++ {
			for z := 0; z <= zSize; z++ {
				fmt.Printf("x: %d - y: %d - z: %d\n", x, y, z)
			}
		}
	}
	for x := 0; x <= 10; x++ {

		if x == 6 {
			break
		}

		fmt.Println(x)
	}
	for x := 0; x <= 10; x++ {

		if x == 6 {
			continue
		}

		fmt.Println(x)
	}
}
