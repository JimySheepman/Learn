package kata

func Well(x []string) string {
	// Your solution!
	count := 0
	for _, v := range x {
		if v == "good" {
			count++
		}
	}
	if count == 0 {
		return "Fail!"
	} else if count == 1 || count == 2 {
		return "Publish!"
	}
	return "I smell a series!"
}
