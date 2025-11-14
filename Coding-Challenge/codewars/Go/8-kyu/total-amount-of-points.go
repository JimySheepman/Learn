package kata

import "strings"

func Points(games []string) int {
	// your code here!
	count := 0
	for _, v := range games {
		st := strings.Split(v, "")
		if st[0] > st[2] {
			count += 3
		} else if st[0] == st[2] {
			count += 1
		}
	}
	return count
}
