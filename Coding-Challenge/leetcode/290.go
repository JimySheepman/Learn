func wordPattern(pattern string, str string) bool {
	list := strings.Split(str, " ")

	dict1 := make(map[byte]string)
	dict2 := make(map[string]byte)

	if len(pattern) != len(list) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if v, ok := dict1[pattern[i]]; !ok {
			dict1[pattern[i]] = list[i]
		} else {
			if v != list[i] {
				return false
			}
		}

		if w, ok := dict2[list[i]]; !ok {
			dict2[list[i]] = pattern[i]
		} else {
			if w != pattern[i] {
				return false
			}
		}
	}

	return true
}
