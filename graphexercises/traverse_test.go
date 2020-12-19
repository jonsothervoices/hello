package graphexercises

import (
	"sort"
	"testing"
)

func TestTraverse(t *testing.T) {
	var tests = []struct {
		d        []int
		expected []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{[]int{7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{[]int{0, 5, 3, 2, 4, 1, 6}, []int{0, 1, 2, 3, 4, 5, 6}},
		{[]int{1, 0, 2}, []int{0, 1, 2}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{0}, []int{0}},
		{[]int{}, []int{}},
	}
	for i, datest := range tests {
		bt := newBST(datest.d, nil)
		actual := bt.traverse()
		sort.Slice(actual, func(i, j int) bool {
			return actual[i] < actual[j]
		})
		if len(datest.expected) != len(actual) {
			t.Errorf("%v: actual length %v, expected length %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range datest.expected {
			if v != actual[j] {
				t.Errorf("%v: actual[%v] %v, expected %v", i, j, actual[j], v)
			}
		}
	}
}
