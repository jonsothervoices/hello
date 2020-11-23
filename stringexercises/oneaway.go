package stringexercises

import "math"

//oneAway has O(n)
func oneAway(a, b string) bool {
	if math.Abs(float64(len(b)-len(a))) > 1 {
		return false
	}
	if len(a) == len(b) {
		return replaceCheck(a, b)
	}
	if len(a)+1 == len(b) {
		return insertCheck(a, b)
	}
	return insertCheck(b, a)
}

func replaceCheck(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return len(a) == i+1 || a[i+1:] == b[i+1:]
		}
	}
	return true
}

func insertCheck(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return a[i:] == b[i+1:]
		}
	}
	return true
}
