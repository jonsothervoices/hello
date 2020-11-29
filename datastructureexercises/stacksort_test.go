package datastructureexercises

import (
	"strings"
	"testing"
)

func intComparer(a, b interface{}) int {
	if a == nil || b == nil {
		return 0
	}
	aint := a.(int)
	bint := b.(int)
	if aint < bint {
		return -1
	}
	if aint == bint {
		return 0
	}
	return 1
}

func strComparer(a, b interface{}) int {
	if a == nil || b == nil {
		return 0
	}
	return strings.Compare(a.(string), b.(string))
}

func TestSort(t *testing.T) {
	var tests = []struct {
		a        *stack
		b        comparer
		expected *stack
	}{
		{
			a:        newStack([]interface{}{2, 5, 1, 7, 3}),
			b:        intComparer,
			expected: newStack([]interface{}{7, 5, 3, 2, 1}),
		},
		{
			a:        newStack([]interface{}{1, 5, 4, 3, 2}),
			b:        intComparer,
			expected: newStack([]interface{}{5, 4, 3, 2, 1}),
		},
		{
			a:        newStack([]interface{}{1, 1}),
			b:        intComparer,
			expected: newStack([]interface{}{1, 1}),
		},
		{
			a:        newStack([]interface{}{}),
			b:        intComparer,
			expected: newStack([]interface{}{}),
		},
		{
			a:        newStack([]interface{}{"b", "c", "a", "A", "r"}),
			b:        strComparer,
			expected: newStack([]interface{}{"r", "c", "b", "a", "A"}),
		},
	}
	for i, datest := range tests {
		actual := stackSorter(datest.a, datest.b)
		if actual.len() != datest.expected.len() {
			t.Logf("%v: actual %v, expected %v", i, actual, datest.expected)
			t.Fail()
		}
		for actual.len() != 0 {
			a := actual.pop()
			e := datest.expected.pop()
			if a != e {
				t.Errorf("%v: actual %v, expected %v", i, a, e)
			}
		}
	}
}
