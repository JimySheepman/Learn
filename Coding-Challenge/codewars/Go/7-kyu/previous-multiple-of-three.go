package kata

func PrevMultOfThree(n int) interface{} {
	// write your code here
	// your function should return an int or a nil
	i := n
	for i > 0 {
		if i%3 == 0 && i != 0 {
			return i
		}
		i = i / 10
	}
	return nil
}
