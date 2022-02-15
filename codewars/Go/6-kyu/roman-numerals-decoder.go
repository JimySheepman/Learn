package kata

var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func Decode(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Decode(roman[2:])
	}
	return first + Decode(roman[1:])
}
