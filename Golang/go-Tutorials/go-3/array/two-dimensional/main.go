package main

import "fmt"

func main() {
	threeD := [2][2][2]string{}
	threeD[0][0][0] = "element 0,0,0"
	threeD[0][1][0] = "element 0,1,0"
	threeD[0][0][1] = "element 0,1,1"
	fmt.Println(threeD)
}
