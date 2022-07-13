package kata

func Maps(x []int) []int {
	for _, v := range x {
		x = append(x, v*2)
	}
	return x[len(x)/2:]
}

/*
func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, v := range x {
	  result[i] = v + v
	}
	return result
  }
*/
