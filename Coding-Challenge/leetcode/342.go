func isPowerOfFour(num int) bool {
	i := 1
	for i < num {
		i *= 4
	}
	return i == num
}