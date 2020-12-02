package stringexercises

import "testing"

func TestStringCompress(t *testing.T) {
	var tests = []struct {
		a        string
		expected string
	}{
		{"ssssaaabcccddee", "s4a3b1c3d2e2"},
		{"ggabbcdde", "ggabbcdde"},
		{"ggg", "g3"},
		{"", ""},
		{"a", "a"},
		{"aabbcccddeeaa", "a2b2c3d2e2a2"},
		{"abbcccdddde", "a1b2c3d4e1"},
	}
	for i, datest := range tests {
		actual := stringCompress(datest.a)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
