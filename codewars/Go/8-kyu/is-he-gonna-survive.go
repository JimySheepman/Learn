package kata

func Hero(bullets, dragons int) bool {
	if bullets-(dragons*2) >= 0 {
		return true
	} else {
		return false
	}
}
