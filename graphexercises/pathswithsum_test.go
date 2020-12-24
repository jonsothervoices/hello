package graphexercises

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	var tests = []struct {
		d        []int
		n        int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, 3},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 8, 3},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 11, 1},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 13, 2},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, 0},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7}, 11, 1},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7}, 6, 2},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7}, 7, 3},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7}, 10, 1},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7}, 18, 0},
		// {[]int{0, 1}, 1, 2},
		// {[]int{0, 1}, 0, 1},
		// {[]int{0, 1}, 2, 0},
		// {[]int{0}, 0, 1},
		// {[]int{0}, 1, 0},
		// {[]int{}, 5, 0},
	}
	for i, datest := range tests {
		b := newBST(datest.d, nil)
		// tester, _ := json.Marshal(b)
		// fmt.Println(string(tester))
		actual := b.pathSum(datest.n)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
