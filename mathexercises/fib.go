package mathexercises

//nth fib number without caching
func uncacheNthFib(n int) int {
	switch {
	case n <= 1:
		{
			return 0
		}
	case n == 2:
		{
			return 1
		}
	}
	return uncacheNthFib(n-2) + uncacheNthFib(n-1)
}

//nth fib number with caching
func cacheNthFib(n int, m map[int]int) int {
	if m == nil {
		m = make(map[int]int)
	}
	if v, ok := m[n]; ok {
		return v
	}
	switch {
	case n <= 1:
		{
			return 0
		}
	case n == 2:
		{
			return 1
		}
	}
	m[n] = cacheNthFib(n-1, m) + cacheNthFib(n-2, m)
	return m[n]
}

//nth fib number with iteration

func itNthFib(n int) int {
	ret := 0
	prev := 0
	current := 1
	for i := 1; i < n; i++ {
		ret = current + prev
		current, prev = ret, current
	}
	return prev
}
