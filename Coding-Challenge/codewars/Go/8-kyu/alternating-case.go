package kata

import "strings"

func ToAlternatingCase(str string) string {
	n := ""
	for _, v := range str {
		if v >= 65 && v <= 90 {
			n += strings.ToLower(string(v))
		} else if v >= 97 && v <= 122 {
			n += strings.ToUpper(string(v))
		} else {
			n += string(v)
		}
	}
	return n
}
