package main

import "testing"

func TestUniqueWordCounter(t *testing.T) {
	var tests = []struct {
		a        []string
		expected int
	}{
		{
			[]string{"to", "be", "or", "not", "to", "be", "that", "is", "the", "question"},
			8,
		},
		{
			[]string{"the", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"},
			8,
		},
		{
			[]string{"the", "rain", "in", "Spain", "falls", "mainly", "on", "the", "yellow", "plain"},
			9,
		},
		{
			[]string{"oh", "captain", "my", "captain"},
			3,
		},
		{
			[]string{""},
			0,
		},
		{
			[]string{},
			0,
		},
		{
			[]string{"", "dont", "count", "", "the", "empty", "strings"},
			5,
		},
	}

	for i, ct := range tests {
		actual := uniqueWordCounter(ct.a)
		if ct.expected != actual {
			t.Errorf("%v: have %v, want %v", i, actual, ct.expected)
		}
	}
}
