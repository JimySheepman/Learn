func earliestFullBloom(plantTime []int, growTime []int) int {
	indexes := make([]int, len(plantTime))

	for i := range indexes {
		indexes[i] = i
	}

	sort.Slice(indexes, func(i, j int) bool {
		if growTime[indexes[i]] != growTime[indexes[j]] {
			return growTime[indexes[i]] > growTime[indexes[j]]
		}

		return plantTime[indexes[i]] > plantTime[indexes[j]]
	})

	now, result := 0, 0

	for _, i := range indexes {
		result = max(result, now+plantTime[i]+growTime[i])
		now += plantTime[i]
	}

	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
