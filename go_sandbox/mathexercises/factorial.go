package mathexercises

func factorial(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	if x > 0 {
		return x * factorial(x-1)
	}
	return x * factorial(x+1)
}
