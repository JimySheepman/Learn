func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	result := [][]int{{1}}
	helper(2, numRows, &result)
	fmt.Println(result)
	return result
}

func helper(n int, numRows int, result *[][]int) {
	if n > numRows {
		return
	}

	level := make([]int, n)
	level[0], level[n-1] = 1, 1
	for i := 1; i < n-1; i++ {
		level[i] = (*result)[len((*result))-1][i-1] + (*result)[len((*result))-1][i]
	}
	*result = append(*result, level)
	helper(n+1, numRows, result)
}
