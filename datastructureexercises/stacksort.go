package datastructureexercises

import "strings"

type Comparer func(a, b interface{}) int

func IntComparer(a, b interface{}) int {
	if a == nil || b == nil {
		return 0
	}
	aint := a.(int)
	bint := b.(int)
	if aint < bint {
		return -1
	}
	if aint == bint {
		return 0
	}
	return 1
}

func StrComparer(a, b interface{}) int {
	if a == nil || b == nil {
		return -2
	}
	return strings.Compare(a.(string), b.(string))
}

//Write a program to sort a stack such that the smallest items are on the top.
//stackSorter is O(n^2)
func stackSorter(s *stack, srt Comparer) (final stack) {
	for s.len() > 0 {
		vplace := s.pop()
		for srt(vplace, final.peek()) > 0 {
			s.push(final.pop())
		}
		final.push(vplace)
	}
	return
}
