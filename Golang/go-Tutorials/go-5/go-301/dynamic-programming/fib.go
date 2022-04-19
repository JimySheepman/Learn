package main

func main() {
	n := 49
	memo := make([]int, n)
	println(Fib2(n, memo))
	println(Fib1(n))
	println(Fib3(n, make([]int, n+1)))
}

func Fib1(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

func Fib2(n int, memo []int) int {
	if n <= 1 {
		return 1
	}
	if memo[n-1] != 0 {
		return memo[n-1]
	}
	res := Fib1(n-1) + Fib1(n-2)
	memo[n-1] = res
	return res
}

func Fib3(n int, up []int) int {
	if n <= 1 {
		return 1
	}
	up[0] = 1
	up[1] = 1
	for i := 2; i <= n; i++ {
		up[i] = up[i-1] + up[i-2]
	}
	return up[n]
}
