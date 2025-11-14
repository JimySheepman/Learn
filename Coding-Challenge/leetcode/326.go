func isPowerOfThree(n int) bool {
	i := 1
	for i < n {
		i *= 3
	}
	return i == n
}