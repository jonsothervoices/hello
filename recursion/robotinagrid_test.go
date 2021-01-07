package recursion

import (
	"testing"
)

func TestRobotGrid(t *testing.T) {
	var tests = []struct {
		m        [][]bool //[]r[]{c1,c2,c3}
		expected []string
		e        bool
	}{
		{[][]bool{
			[]bool{true, false, true},
			[]bool{true, false, false},
			[]bool{true, true, true},
		},
			[]string{"d", "d", "r", "r"}, false},
		{[][]bool{
			[]bool{true, true, true},
			[]bool{true, true, false},
			[]bool{false, true, true},
		},
			[]string{"r", "d", "d", "r"}, false},
		{[][]bool{
			[]bool{true, true, true},
			[]bool{true, true, false},
			[]bool{false, true, false},
		},
			nil, false},
	}
	for i, datest := range tests {
		actual, err := robotGrid(datest.m)
		if err != nil && actual != nil {
			t.Errorf("%v: actual %v, expected error %v", i, actual, err)
			t.FailNow()
		}
		if err == nil && actual == nil {
			t.Errorf("%v: actual nil, expected path %v", i, datest.expected)
			t.FailNow()
		}
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range actual {
			if v != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, v, datest.expected[j])
			}
		}
	}
}
