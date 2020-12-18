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
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 3, 0, 2},
		// {[]int{0, 1, 2, 3, 4, 5, 6}, `{"Data":3,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}},"Right":{"Data":5,"Left":{"Data":4,"Left":null,"Right":null},"Right":{"Data":6,"Left":null,"Right":null}}}`},
		// {[]int{0, 1, 2}, `{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}}`},
		// {[]int{0, 1}, `{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null}`},
		// {[]int{0}, `{"Data":0,"Left":null,"Right":null}`},
		// {[]int{}, "null"},
	}
	for i, datest := range tests {
		bt := newBST(datest.d, nil)
		actual := bt.commonAncestor(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
