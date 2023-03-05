func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		if _, ok := graph[arr[i]]; !ok {
			graph[arr[i]] = []int{i}
		} else {
			graph[arr[i]] = append(graph[arr[i]], i)
		}
	}
	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0}
	step := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == n-1 {
				return step
			}
			nexts := append(graph[arr[curr]], curr-1)
			nexts = append(nexts, curr+1)

			for _, next := range nexts {
				if next >= 0 && next < n && !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}

			delete(graph, arr[curr])
		}
		step++
	}

	return -1
}

