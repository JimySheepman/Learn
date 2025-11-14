func minSetSize(arr []int) int {
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
	}
	s := [][2]int{}
	for k, v := range m {
		s = append(s, [2]int{k, v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][1] > s[j][1]
	})
	current := 0
	half := len(arr)/2 + len(arr)%2
	for i, v := range s {
		current += v[1]
		if current >= half {
			return i + 1
		}
	}
	return 0
}