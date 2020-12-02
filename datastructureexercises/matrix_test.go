package datastructureexercises

import "testing"

func TestStringCompress(t *testing.T) {
	var tests = []struct {
		a        matrix
		expected matrix
	}{
		{matrix{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
			matrix{{"c", "f", "i"}, {"b", "e", "h"}, {"a", "d", "g"}}},
		// {matrix{{"a", "a"}, {"b", "b"}},
		// 	matrix{{"a", "b"}, {"a", "b"}}},
		// {matrix{{"a"}},
		// 	matrix{{"a"}}},
		// {matrix{{}},
		// 	matrix{{}}},
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
