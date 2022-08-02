func kthSmallest(matrix [][]int, k int) int {
	var lst []int
	for _, v := range matrix {
		lst = append(lst, v...)
	}
	sort.Ints(lst)
	return lst[k-1]
}
