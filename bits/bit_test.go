package bits

import (
	"math"
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

func TestFlipToWin(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{5, 3},
		{15, 4},
		{16, 2},
		{17, 2},
		{30, 5},
		{0, 1},
		{11, 4},
		{19, 3},
	}
	for i, datest := range tests {
		actual := flipToWin(datest.n)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestZeroNextToOne(t *testing.T) {
	var tests = []struct {
		s        string
		i        int
		expected bool
	}{
		{"00111011", 5, true},
		{"00111011", 1, true},
		{"00111011", 0, false},
		{"0100111011", 0, true},
		{"001110110", 8, true},
		{"0", 0, true},
		{"00", 0, false},
		{"010", 2, true},
	}
	for i, datest := range tests {
		actual := zeroNextToOne(datest.s, datest.i)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestLSum(t *testing.T) {
	var tests = []struct {
		s        string
		i        int
		expected int
	}{
		{"00111011", 5, 3},
		{"00111011", 3, 1},
		{"00111011", 6, 0},
		{"0100111011", 9, 1},
		{"001110110", 0, 0},
		{"0", 0, 0},
		{"00", 1, 0},
		{"010", 2, 1},
	}
	for i, datest := range tests {
		actual := lSum(datest.s, datest.i)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestRSum(t *testing.T) {
	var tests = []struct {
		s        string
		i        int
		expected int
	}{
		{"00111011", 1, 3},
		{"00111011", 0, 0},
		{"00111011", 7, 0},
		{"0100111011", 0, 1},
		{"001110110", 1, 3},
		{"0", 0, 0},
		{"00", 1, 0},
		{"010", 0, 1},
	}
	for i, datest := range tests {
		actual := rSum(datest.s, datest.i)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestNextLNum(t *testing.T) {
	var tests = []struct {
		n        int64
		expected int64
	}{
		{9, 10},    //1001, 1010
		{59, 61},   //00111011, 00111101
		{315, 317}, //0100111011, 0100111101
		{64, 128},  //01000000, 10000000
		{63, 95},   //00111111, 01011111
		{1, 2},     //00000001, 00000010
		{0, 0},     //00000000, 00000000
		{4294967295, 6442450943},
		{math.MaxInt64, 0},
		{math.MaxInt64 - 1, 0},
	}
	for i, datest := range tests {
		actual, err := nextLNum(datest.n)
		if err != nil && datest.expected != 0 {
			t.Errorf("%v: actual %v, expected %v", i, actual, err)
			t.FailNow()
		}
		if err == nil && datest.expected == 0 {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
			t.FailNow()
		}
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestNextSNum(t *testing.T) {
	var tests = []struct {
		n        int64
		expected int64
	}{
		{9, 5},     //1001, 0101
		{59, 55},   //00111011, 00110111
		{315, 311}, //0100111011, 0100110111
		{64, 32},   //01000000, 00100000
		{63, 0},    //00111111, 00111111
		{1, 0},     //00000001, 00000001
		{0, 0},     //00000000, 00000000
		{2, 1},
		{4294967295, 0}, //11111111111111111111111111111111
		{1022, 1021},
	}
	for i, datest := range tests {
		actual, err := nextSNum(datest.n)
		if err != nil && datest.expected != 0 {
			t.Errorf("%v: actual %v, expected %v", i, actual, err)
			t.FailNow()
		}
		if err == nil && datest.expected == 0 {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
			t.FailNow()
		}
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestConv(t *testing.T) {
	var tests = []struct {
		a        int64
		b        int64
		expected int
	}{
		{9, 5, 2},   //1001, 0101
		{33, 63, 4}, //100001, 111111
		{33, 59, 3}, //100001, 111011
		{12, 12, 0},
		{1, 0, 1},
	}
	for i, datest := range tests {
		actual := conv(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestPairSwap(t *testing.T) {
	var tests = []struct {
		a        int64
		expected int64
	}{
		{9, 6}, //1001
		// {33, 18},   //100001
		// {170, 85},  //10101010, 01010101
		// {366, 669}, //0101101110, 1010011101
		// {1, 2},
		// {2, 1},
	}
	for i, datest := range tests {
		actual := pairSwap(datest.a)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
