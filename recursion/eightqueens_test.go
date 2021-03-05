package recursion

import (
	"fmt"
	"testing"
)

func TestEightQueens(t *testing.T) {
	actual := eightQueens(board{}, 0)
	if len(actual) != 92 {
		t.Errorf("actual %v, expected 92", len(actual))
	}
	fmt.Println(actual)
}

func TestIsAllowed(t *testing.T) {
	var tests = []struct {
		b        board
		i        int
		j        int
		expected bool
	}{
		{board{128, 0, 0, 0, 0, 0, 0, 0}, 1, 0, false},
		{board{128, 64, 32, 16, 8, 4, 2, 1}, 0, 1, false},
		{board{128, 0, 0, 0, 0, 0, 0, 0}, 0, 1, false},
		{board{0, 0, 64, 0, 0, 0, 0, 0}, 0, 1, false},
		{board{0, 0, 16, 0, 0, 0, 0, 0}, 0, 1, false},
		{board{0, 0, 0, 128, 0, 0, 0, 0}, 0, 3, false},
		{board{0, 0, 0, 0, 64, 0, 0, 0}, 0, 2, true},
		{board{0, 0, 0, 0, 0, 0, 0, 0}, 7, 7, true},
		{board{255, 255, 255, 255, 255, 255, 255, 255}, 1, 1, false},
	}
	for i, datest := range tests {
		actual := datest.b.isAllowed(datest.i, datest.j)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestSum(t *testing.T) {
	var tests = []struct {
		b        board
		expected uint8
	}{
		{board{128, 64, 32, 16, 8, 4, 2, 1}, 255},
		{board{128, 0, 0, 0, 0, 0, 0, 0}, 128},
		{board{0, 0, 64, 0, 0, 0, 0, 1}, 65},
		{board{0, 0, 16, 0, 5, 7, 0, 0}, 28},
		{board{0, 0, 0, 0, 0, 0, 0, 0}, 0},
	}
	for i, datest := range tests {
		actual := datest.b.sum()
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
