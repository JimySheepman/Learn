package kata

import "strings"

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str) == str
}
