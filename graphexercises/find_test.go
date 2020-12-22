package graphexercises

import (
	"testing"
)

func TestFind(t *testing.T) {
	var tests = []struct {
		d        []int
		n        int
		expected bool
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 11, false},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, -3, false},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, true},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 100, false},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 3, true},
		{[]int{0, 1, 2}, 0, true},
		{[]int{0, 1}, 3, false},
		{[]int{0}, 0, true},
		{[]int{}, 5, false},
	}
	for i, datest := range tests {
		g := newBST(datest.d, nil)
		// tester, _ := json.Marshal(g)
		// fmt.Println(string(tester))
		actual := g.find(datest.n)
		// str, _ := json.Marshal(node)
		// fmt.Println(string(str))
		if actual == nil && datest.expected {
			t.Errorf("%v: actual %v, expected to find %v", i, actual, datest.n)
			t.FailNow()
		}
		if actual != nil && !datest.expected {
			t.Errorf("%v: actual %v was found, expected nothing", i, actual.Data)
		}
		if actual != nil && actual.Data != datest.n {
			t.Errorf("%v: actual %v, expected %v", i, actual.Data, datest.expected)
		}
	}
}
