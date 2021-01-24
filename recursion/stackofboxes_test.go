package recursion

import (
	"testing"
)

func TestStackOfBoxes(t *testing.T) {
	var tests = []struct {
		s        boxGroup
		expected int
		count    int
	}{
		{boxGroup{box{1, 1, 1}, box{2, 2, 2}, box{3, 3, 3}, box{4, 4, 4}}, 10, 4},
		{boxGroup{box{2, 3, 4}, box{1, 2, 3}, box{4, 5, 6}, box{3, 4, 5}}, 10, 4},
		{boxGroup{box{7, 7, 7}, box{2, 5, 5}, box{1, 4, 4}, box{3, 3, 3}, box{2, 2, 2}, box{1, 1, 1}}, 13, 4},
		{boxGroup{box{5, 2, 2}, box{2, 4, 4}, box{1, 1, 1}}, 6, 2},
		{boxGroup{box{6, 1, 1}, box{5, 2, 2}, box{2, 4, 4}, box{2, 1, 1}}, 7, 2},
		{boxGroup{box{8, 8, 8}, box{7, 8, 9}, box{6, 8, 10}, box{5, 8, 11}}, 8, 1},
		{boxGroup{box{20, 8, 8}, box{7, 10, 9}, box{6, 9, 8}, box{5, 8, 7}}, 20, 1},
	}
	for i, datest := range tests {
		actual := datest.s.stack()
		if len(actual) != datest.count {
			t.Errorf("%v: actual length %v, expected %v", i, len(actual), datest.count)
		}
		hActual := actual.height()
		if hActual != datest.expected {
			t.Errorf("%v: actual height %v, expected %v", i, hActual, datest.expected)
		}
	}
}
