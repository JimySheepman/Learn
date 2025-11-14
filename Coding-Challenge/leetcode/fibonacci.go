package main

func fibLoop(n int) int {
	if n <= 1 {
		return n
	}
	return fibLoop(n-1) + fibLoop(n-2)
}

func fibRecursive(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func fibHMap(n int) int {
	// n limit max
	m := map[int]int{
		0: 0, 1: 1,
		2: 1, 3: 2,
		4: 3, 5: 5,
		6: 8, 7: 13,
		8: 21, 9: 34,
		10: 55, 11: 89,
		12: 144, 13: 233,
		14: 377, 15: 610,
		16: 987, 17: 1597,
		18: 2584, 19: 4181,
		20: 6765, 21: 10946,
		22: 17711, 23: 28657,
		24: 46368, 25: 75025,
		26: 121393, 27: 196418,
		28: 317811, 29: 514229,
		30: 832040,
	}

	return m[n]
}

