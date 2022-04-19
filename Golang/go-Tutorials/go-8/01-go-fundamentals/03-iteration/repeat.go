package iteration

const repatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repatCount; i++ {
		repeated += character
	}
	return repeated
}
