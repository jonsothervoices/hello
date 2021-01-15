package recursion

import (
	"sort"
	"testing"
)

func TestNonUniquePerm(t *testing.T) {
	var tests = []struct {
		a        string
		expected []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"aa", []string{"aa"}},
		{"aaa", []string{"aaa"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aac", []string{"aac", "aca", "caa"}},
		{"abac", []string{"aabc", "aacb", "abac", "abca", "acab", "acba", "baac", "baca", "bcaa", "caab", "caba", "cbaa"}},
		{"abcd", []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}},
	}
	for i, datest := range tests {
		actual := nonUniquePerm(datest.a)
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
