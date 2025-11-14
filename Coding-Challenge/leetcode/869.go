func reorderedPowerOf2(n int) bool {
	ns := strconv.Itoa(n)
	nm := map[rune]int{}

	for _, nsb := range ns {
		nm[nsb]++
	}

	ceil := 10
	for ceil < n {
		ceil *= 10
	}

	for i := 0; 1<<i <= ceil; i++ {
		if 1<<i < ceil/10 {
			continue
		}

		s := strconv.Itoa(1 << i)

		for j := '0'; j <= '9'; j++ {
			if count(s, rune(j)) != nm[j] {
				break
			}
			if j == '9' {
				return true
			}
		}
	}

	return false

}

func count(s string, b rune) int {
	count := 0
	for _, sb := range s {
		if sb == b {
			count++
		}
	}
	return count
}
