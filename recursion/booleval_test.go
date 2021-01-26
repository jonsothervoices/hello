package recursion

import (
	"testing"
)

func TestBoolEval(t *testing.T) {
	var tests = []struct {
		s        string
		result   bool
		expected int
	}{
		// // testing XOR logic
		{"1^0", false, 0},
		{"1^0", true, 1},
		{"1^1", true, 0},
		{"1^1", false, 1},
		{"0^0", true, 0},
		{"0^0", false, 1},
		// // testing OR logic
		{"1|0", true, 1},
		{"1|0", false, 0},
		{"1|1", true, 1},
		{"1|1", false, 0},
		{"0|0", true, 0},
		{"0|0", false, 1},
		//testing AND logic
		{"1&0", true, 0},
		{"1&0", false, 1},
		{"1&1", true, 1},
		{"1&1", false, 0},
		{"0&0", false, 1},
		{"0&0", true, 0},
		//test cases
		{"1&1|0", true, 2},
		{"1&0&1|0", false, 5},
		{"1&0&1|0", true, 0},
		{"1^0|0|1", false, 2},
		{"0&0&0&1^1|0", true, 10},
		{"1&1^0|0", true, 5},
		{"1&1^0|0", false, 0},
		{"1&0^0|0", false, 5},
		{"1&0^0|0", true, 0},
		{"1", true, 1},
		{"1", false, 0},
		{"0", true, 0},
		{"0", false, 1},
	}
	for i, datest := range tests {
		actual := boolEval(datest.s, datest.result)
		if actual != datest.expected {
			t.Errorf("%v: actual  %v, expected %v", i, actual, datest.expected)
		}
	}
}
