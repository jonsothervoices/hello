package datastructureexercises

import (
	"testing"
)

func numStackCheck(in int) int {
	if in == 0 {
		return 0
	}
	if in%5 == 0 {
		return in / 5
	}
	return in/5 + 1
}
func TestSuperstackPop(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		lenExp   int
		expected interface{}
	}{
		{[]interface{}{"a", "b", "c", "d", "e"}, 4, "e"},
		{[]interface{}{"a", "b", "c", "d", "e", "f", "g"}, 6, "g"},
		{[]interface{}{"a", "b", "c", "d", "e", "f"}, 5, "f"},
		{[]interface{}{"a"}, 0, "a"},
		{[]interface{}{""}, 0, ""},
		{[]interface{}{}, 0, nil},
	}
	for i, datest := range tests {
		//each iteration, we build the superstack from test data (d)
		sa := newSuperStack(datest.a)
		actual := sa.pop()
		//did we pop the right thing?
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		//did we properly alter the length of our superstack?
		if sa.len() != datest.lenExp {
			t.Errorf("%v: length %v, expected %v", i, sa.len(), datest.lenExp)
		}
		//did we remove the empty stack at the end if we have one?
		if len(datest.a) == 0 {
			if sa.numStacks() != 0 {
				t.Errorf("%v: numStacks %v, expected %v", i, sa.numStacks(), 0)
			}
		} else if sa.numStacks() != numStackCheck(len(datest.a)-1) {
			t.Errorf("%v: numStacks %v, expected %v", i, sa.numStacks(), numStackCheck(len(datest.a)-1))
		}
	}
}

func TestSuperstackPush(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		b        string
		expected string
		lenExp   int
	}{
		{[]interface{}{"a", "b", "c", "d", "e"}, "f", "f", 6},
		{[]interface{}{"a", "b", "c", "d"}, "e", "e", 5},
		{[]interface{}{"a"}, "b", "b", 2},
		{[]interface{}{""}, "", "", 2},
		{[]interface{}{}, "", "", 1},
	}
	for i, datest := range tests {
		//each iteration, we build the superstack from test data (d)
		sa := newSuperStack(datest.a)
		sa.push(datest.b)
		actual := sa.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if sa.numStacks() != numStackCheck(len(datest.a)+1) {
			t.Errorf("%v: numStacks %v, expected %v", i, sa.numStacks(), numStackCheck(len(datest.a)+1))
		}
	}
}

func TestSuperstackPeek(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		expected interface{}
	}{
		{[]interface{}{"a", "b", "c", "d", "e", "f", "g"}, "g"},
		{[]interface{}{"a", "b", "c", "d"}, "d"},
		{[]interface{}{"a"}, "a"},
		{[]interface{}{""}, ""},
		{[]interface{}{}, nil},
	}
	for i, datest := range tests {
		sa := newSuperStack(datest.a)
		actual := sa.peek()
		actual2 := sa.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if actual != actual2 {
			t.Errorf("%v: stack altered from %v to %v", i, actual, actual2)
		}
		if sa.numStacks() != numStackCheck(len(datest.a)) {
			t.Errorf("%v: numStacks %v, expected %v", i, sa.numStacks(), numStackCheck(len(datest.a)))
		}
	}
}
