package kata

import "strings"

func NoSpace(word string) string {
	return strings.Join(strings.Split(word, " "), "")
	// return strings.ReplaceAll(word, " ", "")
}
