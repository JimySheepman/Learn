package kata

import "fmt"

func countSheep(num int) string {
	// Your code here!
	res := ""
	for i := 0; i < num; i++ {
		res += fmt.Sprintf("%d sheep...", i+1)
	}
	return res
}
