func lengthOfLastWord(s string) int {
	pos := len(s) - 1
	for ; pos >= 0; pos-- {
		if !unicode.IsSpace(rune(s[pos])) {
			break
		}
	}
	s = s[:pos + 1]
	if len(s) == 0 {
		return 0
	}

	for ; pos >= 0; pos-- {
		if unicode.IsSpace(rune(s[pos])) {
			break
		}
	}
	return len(s[pos+1:])
}
