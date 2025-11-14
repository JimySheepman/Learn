func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	var a, b int = 1, 2
	for i := 3; i <= n; i++ {
		b += a
		a = b - a
	}
	return b
}