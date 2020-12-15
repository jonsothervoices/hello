package graphexercises

import (
	"testing"
)

func TestSuccessor(t *testing.T) {
	var tests = []struct {
		d        []int
		n        int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, 2},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, 3},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 4},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, 5},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 6},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 6, 7},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, 8},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 8, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 0},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 3, 4},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 0, 1},
		{[]int{0, 1, 2}, 1, 2},
		{[]int{0, 1}, 1, 0},
		{[]int{0}, 0, 0},
		{[]int{}, 5, 0},
	}
	for i, datest := range tests {
		g := newBST(datest.d, nil)
		// tester, _ := json.Marshal(g)
		// fmt.Println(string(tester))
		node := g.find(datest.n)
		// str, _ := json.Marshal(node)
		// fmt.Println(string(str))
		actual := node.successor()
		if actual == nil && datest.expected != 0 {
			t.Errorf("%v: no successor", i)
		} else if actual != nil && datest.expected != actual.Data {
			t.Errorf("%v: actual %v, expected %v", i, actual.Data, datest.expected)
		}
	}
}
