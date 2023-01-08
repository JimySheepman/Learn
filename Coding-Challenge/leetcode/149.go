func maxPoints(points [][]int) int {
	result := 0
	sort.Slice(points, func(a, b int) bool {
		return points[a][0] < points[b][0]
	})
	for i := 0; i < len(points); i++ {
		m := make(map[float32]int)
		count := 0
		for j := i + 1; j < len(points); j++ {
			if points[j][0] == points[i][0] {
				count += 1
			} else {
				ratio := float32(points[j][1]-points[i][1]) / float32(points[j][0]-points[i][0])
				m[ratio] += 1
				if m[ratio] > result {
					result = m[ratio]
				}
			}
		}
		if count > result {
			result = count
		}
	}

	return result + 1
}
