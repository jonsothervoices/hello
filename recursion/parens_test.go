package recursion

import (
	"testing"
)

func TestParens(t *testing.T) {
	var tests = []struct {
		s        string
		expected bool
	}{
		{"the (rain) in spain)", false},
		{"falls (mainly)", true},
		{"", true},
		{"(in)", true},
		{"(the", false},
		{"(pl(a)in)", true},
		{"(a)a(c)", true},
		{")a(b(ac)", false},
		{"((fff)(", false},
		{"((())", false},
	}
	for i, datest := range tests {
		actual := parens(datest.s)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
