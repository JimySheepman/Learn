package kata

func Past(h, m, s int) int {
	return (h * 60 * 60 * 1000) + (m * 60 * 1000) + (s * 1000)
}
