package stringexercises

import "testing"

func TestRotationCheck(t *testing.T) {
	var tests = []struct {
		a        string
		b        string
		expected bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"table", "bleta", true},
		{"waterbottle", "erbottlewa", false},
		{"table", "tab", false},
		{"table", "waterbottle", false},
		{"ab", "ba", true},
		{"a", "a", true},
		{"", "", true},
	}
	for i, datest := range tests {
		actual := rotationCheck(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestBetterRotationCheck(t *testing.T) {
	var tests = []struct {
		a        string
		b        string
		expected bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"table", "bleta", true},
		{"waterbottle", "erbottlewa", false},
		{"table", "tab", false},
		{"table", "waterbottle", false},
		{"ab", "ba", true},
		{"a", "a", true},
		{"", "", true},
	}
	for i, datest := range tests {
		actual := betterRotationCheck(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
