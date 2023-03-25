func convertToTitle(columnNumber int) string {
	result := ""
	for 0 < columnNumber {
		tmp := columnNumber % 26
		if tmp == 0 {
			tmp = 26
		}
		columnNumber = (columnNumber - tmp) / 26
		result = toChar(tmp) + result
	}
	return result
}

func toChar(i int) string {
	return string('A' - 1 + i)
}

