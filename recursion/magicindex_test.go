package recursion

import (
	"testing"
)

func TestMagicIndex(t *testing.T) {
	var tests = []struct {
		s        []int
		expected int
	}{
		{[]int{1, 1, 7, 11, 20}, 1},
		{[]int{-10, -8, 2, 4, 34}, 2},
		{[]int{-12, -10, 1, 3, 34}, 3},
		{[]int{0, 1, 2, 4, 34}, 2},
		{[]int{-10, -8, 1, 2, 4, 6}, 4},
		{[]int{-10, -8, 0, 4, 34}, -1},
		{[]int{-10}, -1},
		{[]int{0}, 0},
	}
	for i, datest := range tests {
		actual := magicIndex(datest.s)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
