package dynamic

func Fib(n int) int {
	memo := make(map[int]int)

	return fibHelper(n, memo)
}

func fibHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibHelper(n-2, memo) + fibHelper(n-1, memo)

	return memo[n]
}
