package main

import "testing"

func TestFactorial(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{0, 1},
		{18, 6402373705728000},
		{-10, 3628800},
		{-11, -39916800},
		{-1, -1},
		{1, 1},
	}
	for i, ft := range tests {
		have := factorial(ft.a)
		if ft.want != have {
			t.Errorf("%v: have %v, want %v", i, have, ft.want)
		}
	}
}
