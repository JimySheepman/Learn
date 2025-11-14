package kata

import "fmt"

func lonelyinteger(a []int32) int32 {
	// Write your code here
	count := 0
	var t int32
	for i := 0; i < len(a); i++ {
		fmt.Println("i->", a[i])
		for j := 0; j < len(a); j++ {
			fmt.Println("j->", a[j])
			if i != j {
				if a[i] == a[j] {
					fmt.Println("Ai=Aj")
					count++
				}
			}
		}
		fmt.Println("count = ", count)
		if count == 0 {
			t = int32(a[i])
		}
		count = 0
	}
	return t
}
