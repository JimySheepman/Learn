package kata

func CheckForFactor(base int, factor int) bool {
	if base%factor == 0 {
		return true
	}
	return false
}

/*
func CheckForFactor(base int, factor int) bool {
    return base % factor == 0;
}
*/
