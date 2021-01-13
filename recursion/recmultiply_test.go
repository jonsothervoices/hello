package recursion

import (
	"testing"
)

func TestRecMultiply(t *testing.T) {
	var tests = []struct {
		a        uint
		b        uint
		expected uint
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 45, 45},
		{0, 10, 0},
		{25, 0, 0},
		{0, 10, 0},
		{1, 24, 24},
		{24, 1, 24},
		{2, 4, 8},
		{2, 50, 100},
		{8, 8, 64},
		{1, 1, 1},
		{4, 10000000000, 40000000000},
	}
	for i, datest := range tests {
		actual := recMultiply(datest.a, datest.b)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
