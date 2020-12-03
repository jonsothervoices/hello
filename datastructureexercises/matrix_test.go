package datastructureexercises

import "testing"

func TestMatrixRotate(t *testing.T) {
	var tests = []struct {
		a        matrix
		expected matrix
	}{
		{matrix{{"00", "10", "20", "30"}, {"01", "11", "21", "31"}, {"02", "12", "22", "32"}, {"03", "13", "23", "33"}},
			matrix{{"03", "02", "01", "00"}, {"13", "12", "11", "10"}, {"23", "22", "21", "20"}, {"33", "32", "31", "30"}}},
		{matrix{{"00", "10", "20"}, {"01", "11", "21"}, {"02", "12", "22"}},
			matrix{{"02", "01", "00"}, {"12", "11", "10"}, {"22", "21", "20"}}},
		{matrix{{"a", "a"}, {"b", "b"}},
			matrix{{"b", "a"}, {"b", "a"}}},
		{matrix{{"a"}},
			matrix{{"a"}}},
		{matrix{{}},
			matrix{{}}},
	}
	for i, datest := range tests {
		datest.a.rotate()
		for j, v := range datest.a {
			for k, u := range v {
				if u != (datest.expected)[j][k] {
					t.Errorf("%v: [%v] [%v], expected %v actual %v", i, j, k, (datest.expected)[j][k], u)
				}
			}
		}
	}
}
