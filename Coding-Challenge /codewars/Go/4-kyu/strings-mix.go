package main

import (
	"fmt"
	"sort"
	"strings"
)

func Mix(s1, s2 string) string {
	alphabase := "abcdefghijklmnopqrstuvwxyz"
	result := []string{}
	for _, c := range alphabase {
		nb_s1 := strings.Count(s1, string(c))
		nb_s2 := strings.Count(s2, string(c))
		if nb_s1 > 1 || nb_s2 > 1 {
			if nb_s1 == nb_s2 {
				result = append(result, "=:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 > nb_s2 {
				result = append(result, "1:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 < nb_s2 {
				result = append(result, "2:"+strings.Repeat(string(c), nb_s2))
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) > len(result[j])
	})
	return strings.Join(result, "/")
}

func main() {
	s1 := "mmmmm m nnnnn y&friend&Paul has heavy hats! &"
	s2 := "my frie n d Joh n has ma n y ma n y frie n ds n&"
	res := "1:mmmmmm/=:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

	got := Mix(s1, s2)
	fmt.Println(got)
	if res == got {
		fmt.Println("true")
	}
	fmt.Println("Program Exit!")
}
