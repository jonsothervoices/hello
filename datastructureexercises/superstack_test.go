package datastructureexercises

import "testing"

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
		sa := superstack{}
		for _, d := range datest.a {
			sa.push(d)
		}
		actual := sa.pop()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if sa.len() != datest.lenExp {
			t.Errorf("%v: length %v, expected %v", i, sa.len(), datest.lenExp)
		}
		if sa.numStacks() != (len(datest.a)-1)/5+1 {
			t.Errorf("%v: numStacks %v, expected %v", i, sa.numStacks(), (len(datest.a)-1)/5+1)
		}
	}
}

func TestSuperstackPush(t *testing.T) {
	var tests = []struct {
		a        stack
		b        string
		expected string
		lenExp   int
	}{
		{stack{[]interface{}{"a", "b", "c", "d"}}, "e", "e", 5},
		{stack{[]interface{}{"a"}}, "b", "b", 2},
		{stack{[]interface{}{""}}, "", "", 2},
		{stack{[]interface{}{}}, "", "", 1},
	}
	for i, datest := range tests {
		datest.a.push(datest.b)
		actual := datest.a.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestSuperstackPeek(t *testing.T) {
	var tests = []struct {
		a        stack
		expected string
	}{
		{stack{[]interface{}{"a", "b", "c", "d"}}, "d"},
		{stack{[]interface{}{"a"}}, "a"},
		{stack{[]interface{}{""}}, ""},
	}
	for i, datest := range tests {
		actual := datest.a.peek()
		actual2 := datest.a.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if actual != actual2 {
			t.Errorf("%v: stack altered from %v to %v", i, actual, actual2)
		}
	}
}
