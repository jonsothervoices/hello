package mathexercises

import "testing"

func TestCollatzCounter(t *testing.T) {
	var tests = []struct {
		a        uint
		expected uint
	}{
		{12, 9},
		{9, 19},
		{871, 178},
		{0, 0},
		{1, 0},
	}
	for i, ct := range tests {
		actual := collatzCounter(ct.a)
		if ct.expected != actual {
			t.Errorf("%v: have %v, want %v", i, actual, ct.expected)
		}
	}
}
