package main

func CheckPolidrom(str string) bool {
	s := []byte(str)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s) == str
}
