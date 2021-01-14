package recursion

import (
	"sort"
	"testing"
)

func TestUniquePerm(t *testing.T) {
	var tests = []struct {
		a        string
		expected []string
	}{
		{"", []string{""}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"abcd", []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}},
	}
	for i, datest := range tests {
		actual := uniquePerm(datest.a)
		sort.Strings(actual)
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range actual {
			if v != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, v, datest.expected[j])
			}
		}
	}
}
