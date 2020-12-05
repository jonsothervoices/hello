package datastructureexercises

import "testing"

func TestZeroMatrix(t *testing.T) {
	var tests = []struct {
		a        intmatrix
		expected intmatrix
	}{
		{intmatrix{{0, 10, 20, 30}, {01, 11, 21, 31}, {02, 12, 22, 32}, {03, 13, 23, 33}},
			intmatrix{{0, 0, 0, 0}, {0, 11, 21, 31}, {0, 12, 22, 32}, {0, 13, 23, 33}}},
		{intmatrix{{5, 10, 20}, {01, 0, 21}, {02, 12, 22}},
			intmatrix{{5, 0, 20}, {0, 0, 0}, {02, 0, 22}}},
		{intmatrix{{1, 1}, {1, 0}},
			intmatrix{{1, 0}, {0, 0}}},
		{intmatrix{{0}},
			intmatrix{{0}}},
		{intmatrix{{}},
			intmatrix{{}}},
		{intmatrix{{0, 10, 0, 30}, {0, 11, 0, 31}, {02, 12, 22, 32}, {03, 13, 23, 33}, {04, 14, 24, 34}},
			intmatrix{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 12, 0, 32}, {0, 13, 0, 33}, {0, 14, 0, 34}}},
	}
	for i, datest := range tests {
		datest.a.ZeroMatrix()
		for j, v := range datest.a {
			for k, u := range v {
				if u != (datest.expected)[j][k] {
					t.Errorf("%v: [%v] [%v], expected %v actual %v", i, j, k, (datest.expected)[j][k], u)
				}
			}
		}
	}
}
