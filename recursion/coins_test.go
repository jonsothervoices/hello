package recursion

import (
	"testing"
)

func TestCoins(t *testing.T) {
	var tests = []struct {
		a        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{5, 2},
		{9, 2},
		{10, 4},
		{17, 6},
		{41, 31},
		{25, 13},
		{100, 242},
	}
	for i, datest := range tests {
		actual := coins(datest.a)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
