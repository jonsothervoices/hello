package graphexercises

import (
	"testing"
)

func TestCommonAncestor(t *testing.T) {
	var tests = []struct {
		d        []int
		a        int
		b        int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 3, 8, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 3, 0, 2},
		{[]int{0, 5, 3, 2, 4, 1, 6}, 3, 2, 2},
		{[]int{1, 0, 2}, 1, 2, 0},
		{[]int{0, 1}, 1, 0, 1},
		{[]int{0}, 0, 3, -1},
		{[]int{}, 4, 2, -1},
	}
	for i, datest := range tests {
		bt := newBST(datest.d, nil)
		actual := bt.commonAncestor(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
