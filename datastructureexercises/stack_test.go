package datastructureexercises

import "testing"

func TestPop(t *testing.T) {
	var tests = []struct {
		a        stack
		lenExp   int
		expected interface{}
	}{
		{stack{[]interface{}{"a", "b", "c", "d", "e"}}, 4, "e"},
		{stack{[]interface{}{"a"}}, 0, "a"},
		{stack{[]interface{}{""}}, 0, ""},
		{stack{[]interface{}{}}, 0, nil},
	}

	for i, datest := range tests {
		actual := datest.a.pop()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if datest.a.len() != datest.lenExp {
			t.Errorf("%v: length %v, expected %v", i, len(datest.a.data), datest.lenExp)
		}
	}
}

func TestPush(t *testing.T) {
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

func TestQueuePeek(t *testing.T) {
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

func TestMin(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		expected interface{}
	}{
		{[]interface{}{0, 2, 5, 12, 110}, 0},
		{[]interface{}{110, 72, 50, 12, 11}, 11},
		{[]interface{}{0, 2, 5, 0, 11}, 0},
		{[]interface{}{3, -10, 0, 5, 33}, -10},
		{[]interface{}{4}, 4},
		{[]interface{}{}, nil},
	}

	for i, datest := range tests {
		newM := newMinStack(datest.a, IntComparer)
		actual := newM.findmin()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		intExp, _ := datest.expected.(int)
		newM.push(intExp - 1)
		actual = newM.findmin()
		if intExp-1 != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		newM.pop()
		actual = newM.findmin()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
