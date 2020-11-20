//Given a string, write a function to check if it is a permutation of a palindrome.
//Basically, can the string be rearranged to form a palendrome?

package stringexercises

import "testing"

func TestPalPermute(t *testing.T) {
	var tests = []struct {
		a        string
		expected bool
	}{
		{"No lemons no melon", true},
		{"racecar", true},
		{"it is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", false},
		{"s", true},
		{"", true},
		{"aaa", true},
	}
	for i, datest := range tests {
		actual := PalPermute(datest.a)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
