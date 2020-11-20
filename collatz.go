package main

func collatzCounter(n uint) uint {
	if n <= 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + collatzCounter(n/2)
	}
	return 2 + collatzCounter((3*n+1)/2)
}
