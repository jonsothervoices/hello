package bits

import (
	"testing"
)

func TestGetBit(t *testing.T) {
	var tests = []struct {
		n        int
		b        int
		expected bool
	}{
		{9, 0, true},
		{9, 2, false},
		{20, 2, true},
		{20, 3, false},
		{1, 19, false},
		{-255, 0, true},
		{-255, 4, false},
		{-256, 8, true},
		{-256, 0, false},
	}
	for i, datest := range tests {
		actual := getBit(datest.n, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestSetBit(t *testing.T) {
	var tests = []struct {
		n        int
		b        int
		expected int
	}{
		{9, 1, 11},
		{9, 2, 13},
		{9, 4, 25},
		{20, 1, 22},
		{20, 0, 21},
		{20, 2, 20},
		{1, 19, 524289},
		{0, 5, 32},
		{-255, 3, -247},
		{-255, 4, -239},
	}
	for i, datest := range tests {
		actual := setBit(datest.n, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

//test data be flipped up in herrrr!!!!!!!!!
func TestClearBit(t *testing.T) {
	var tests = []struct {
		expected int
		b        int
		n        int
	}{
		{9, 1, 11},
		{9, 2, 13},
		{9, 4, 25},
		{20, 1, 22},
		{20, 0, 21},
		{16, 2, 20},
		{1, 19, 524289},
		{0, 5, 32},
		{-255, 3, -247},
		{-255, 4, -239},
	}
	for i, datest := range tests {
		actual := clearBit(datest.n, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

//5.1 Insertion test (test data be flipped!!!)
func TestInsertBits(t *testing.T) {
	var tests = []struct {
		m        int
		n        int
		i        int
		expected int
	}{
		{9, 20, 3, 25},
		{9, 2, 5, 38},
		{10, 4, 7, 164},
		{1, 20, 3, 28},
		{7, -255, 5, -199},
	}
	for i, datest := range tests {
		actual := insertBits(datest.n, datest.m, datest.i)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
