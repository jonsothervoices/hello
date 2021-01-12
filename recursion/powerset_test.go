package recursion

import (
	"fmt"
	"testing"
)

func TestPowerSet(t *testing.T) {
	var tests = []struct {
		s        []int
		expected [][]int
	}{
		{[]int{1, 1, 7, 11},
			[][]int{
				[]int{},
				[]int{11},
				[]int{7},
				[]int{7, 11},
				[]int{1},
				[]int{1, 11},
				[]int{1, 7},
				[]int{1, 7, 11}}},

		{[]int{1, 9, 3},
			[][]int{
				[]int{},
				[]int{9},
				[]int{3},
				[]int{3, 9},
				[]int{1},
				[]int{1, 9},
				[]int{1, 3},
				[]int{1, 3, 9}}},

		{[]int{1, 2},
			[][]int{
				[]int{},
				[]int{2},
				[]int{1},
				[]int{1, 2}}},
		{[]int{}, [][]int{[]int{}}},
	}
	for i, datest := range tests {
		actual := powerSet(datest.s)
		fmt.Printf("%v\n", actual)
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected length %v\n", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range datest.expected {
			if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", actual[j]) {
				t.Errorf("%v: %v: actual %v, expected %v", i, j, actual[j], datest.expected[j])
			}
		}
	}
}
