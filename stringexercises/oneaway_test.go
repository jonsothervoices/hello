package stringexercises

import "testing"

func TestOneAwayInsert(t *testing.T) {
	var tests = []struct {
		a        string
		b        string
		expected bool
	}{
		{"No lemons no melon", "No lemons no melon", true},
		{"racecar", "racecars", true},
		{"it is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", "iit is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", true},
		{"s", "sd", true},
		{"s", "sdd", false},
		{"", "d", true},
		{"", "", true},
		{"abcdef", "abcdfeg", false},
	}
	for i, datest := range tests {
		actual := oneAway(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestOneAwayRemove(t *testing.T) {
	var tests = []struct {
		a        string
		b        string
		expected bool
	}{
		{"No lemons no melon", "No lemons no melon", true},
		{"racecar", "racecars", true},
		{"it is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", "iit is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", true},
		{"s", "sd", true},
		{"s", "sdd", false},
		{"", "d", true},
		{"", "", true},
		{"abcdef", "abcdfeg", false},
	}
	for i, datest := range tests {
		actual := oneAway(datest.b, datest.a)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestOneAwayReplace(t *testing.T) {
	var tests = []struct {
		a        string
		b        string
		expected bool
	}{
		{"No lemons no melon", "No lemons no melon", true},
		{"racecar", "racecaR", true},
		{"it is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", "It is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", true},
		{"s", "d", true},
		{"sd", "ds", false},
		{"ab", "cd", false},
		{"abcdef", "fbcdea", false},
		{" ", "d", true},
		{"", "", true},
	}
	for i, datest := range tests {
		actual := oneAway(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
