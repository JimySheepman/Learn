package kata

func GetGrade(a, b, c int) rune {
	average := (a + b + c) / 3
	if 90 <= average && average <= 100 {
		return 'A'
	} else if 80 <= average && average < 90 {
		return 'B'
	} else if 70 <= average && average < 80 {
		return 'C'
	} else if 60 <= average && average < 70 {
		return 'D'
	} else {
		return 'F'
	}

}
