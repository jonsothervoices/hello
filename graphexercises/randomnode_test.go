package graphexercises

import (
	"math/rand"
	"testing"
)

func TestGetRandomNode(t *testing.T) {
	// setup()
	defer teardown()

	var tests = []struct {
		d        []int
		seed     int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 5, 5},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 1, 1},
		{[]int{0, 5, 3, 2, 4, 1, 6}, 2, 3},
		{[]int{1, 0, 2}, 0, 1},
		{[]int{0, 1}, 1, 1},
		{[]int{0}, 0, 0},
		{[]int{}, 1, -1},
	}
	for i, datest := range tests {
		intn = func(i int) int {
			// fmt.Printf("stubbed intN returning %v\n", i)
			return datest.seed
		}
		bt := newBST(datest.d, nil)
		actual, err := bt.getRandomNode()
		if err != nil && datest.expected > -1 {
			t.Errorf("%v: unexpected error %v, expected %v", i, err, datest.expected)
		}
		if err == nil && datest.expected == -1 {
			t.Errorf("%v: expected error, actual %v", i, actual)
		}
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func teardown() {
	intn = rand.Intn
}
