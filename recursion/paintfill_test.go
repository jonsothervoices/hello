package recursion

import (
	"testing"
)

func TestPaintFill(t *testing.T) {
	var tests = []struct {
		image    [][]string
		c        string
		x        int
		y        int
		expected [][]string
	}{
		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#AAAAAA"},
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}},
			"#BBBBBB", 1, 1,
			[][]string{
				[]string{"#FFFFFF", "#BBBBBB", "#FFFFFF"},
				[]string{"#BBBBBB", "#BBBBBB", "#BBBBBB"},
				[]string{"#FFFFFF", "#BBBBBB", "#FFFFFF"}}},

		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#AAAAAA"},
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}},
			"#BBBBBB", 0, 0,
			[][]string{
				[]string{"#BBBBBB", "#AAAAAA", "#FFFFFF"},
				[]string{"#AAAAAA", "#AAAAAA", "#AAAAAA"},
				[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}}},

		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#BBBBBB"},
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}},
			"#BBBBBB", 2, 2,
			[][]string{
				[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
				[]string{"#AAAAAA", "#AAAAAA", "#BBBBBB"},
				[]string{"#FFFFFF", "#AAAAAA", "#BBBBBB"}}},

		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#BBBBBB"},
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}},
			"#BBBBBB", 2, 3,
			[][]string{
				[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
				[]string{"#AAAAAA", "#AAAAAA", "#BBBBBB"},
				[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}}},

		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#BBBBBB"},
			[]string{"#FFFFFF", "#AAAAAA", "#BBBBBB"}},
			"#BBBBBB", 2, 2,
			[][]string{
				[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
				[]string{"#AAAAAA", "#AAAAAA", "#BBBBBB"},
				[]string{"#FFFFFF", "#AAAAAA", "#BBBBBB"}}},

		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#AAAAAA"},
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}},
			"#FFFFFF", 0, 1,
			[][]string{
				[]string{"#FFFFFF", "#FFFFFF", "#FFFFFF"},
				[]string{"#FFFFFF", "#FFFFFF", "#FFFFFF"},
				[]string{"#FFFFFF", "#FFFFFF", "#FFFFFF"}}},

		{[][]string{
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF", "#FFFFFF"},
			[]string{"#AAAAAA", "#AAAAAA", "#AAAAAA"},
			[]string{"#FFFFFF", "#AAAAAA", "#FFFFFF"}},
			"#FFFFFF", 0, 1,
			[][]string{
				[]string{"#FFFFFF", "#FFFFFF", "#FFFFFF", "#FFFFFF"},
				[]string{"#FFFFFF", "#FFFFFF", "#FFFFFF"},
				[]string{"#FFFFFF", "#FFFFFF", "#FFFFFF"}}},

		{[][]string{},
			"#FFFFFF", 0, 1,
			[][]string{}},

		{[][]string{[]string{"#AAAAAA"}},
			"#FFFFFF", 0, 0,
			[][]string{[]string{"#FFFFFF"}}},
	}
	for i, datest := range tests {
		actual := paintFill(datest.image, datest.c, datest.x, datest.y)
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual outer length %v, expected %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range actual {
			if len(v) != len(datest.expected[j]) {
				t.Errorf("%v: actual inner length %v, expected %v", i, len(v), len(datest.expected[j]))
				t.FailNow()
			}
			for k, u := range v {
				if u != datest.expected[j][k] {
					t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
				}
			}
		}
	}
}
