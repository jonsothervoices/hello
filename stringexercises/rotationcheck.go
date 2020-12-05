package stringexercises

import "strings"

//Given 2 strings, write code to check if s2 is a rotation of s1 using only 1 call to strings.Contains()

//Contains(s, substr string) bool

func rotationCheck(a, b string) bool {
	if a == b {
		return true
	}
	for range a {
		a = string(a[len(a)-1]) + a[0:len(a)-1]
		if a == b {
			return true
		}
	}
	return false
}

func betterRotationCheck(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return strings.Contains(a+a, b)
}
