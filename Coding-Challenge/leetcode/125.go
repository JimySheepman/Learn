func isPalindrome(s string) bool {
	var (
		newStr = strings.ToLower(strings.ReplaceAll(s, " ", ""))
		list   = []rune(newStr)
		first  = 0
		last   = len(list) - 1
	)

	for first <= last {
		var (
			left  = list[first]
			right = list[last]
		)

		if left == right {
			first++
			last--
			continue
		} else if !unicode.IsLetter(left) && !unicode.IsLetter(right) && !unicode.IsNumber(left) && !unicode.IsNumber(right) {
			last--
			first++
		} else if !unicode.IsLetter(left) && !unicode.IsNumber(left) {
			first++
		} else if !unicode.IsLetter(right) && !unicode.IsNumber(right) {
			last--
		} else {
			return false
		}
	}

	return true
}