package datastructureexercises

import "testing"

func TestRemove(t *testing.T) {
	var tests = []struct {
		a        queue
		lenExp   int
		expected interface{}
	}{
		{queue{[]interface{}{"a", "b", "c", "d", "e"}}, 4, "a"},
		{queue{[]interface{}{"a"}}, 0, "a"},
		{queue{[]interface{}{""}}, 0, ""},
		{queue{[]interface{}{}}, 0, nil},
	}
	for i, datest := range tests {
		actual := datest.a.remove()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if datest.a.len() != datest.lenExp {
			t.Errorf("%v: length %v, expected %v", i, len(datest.a.data), datest.lenExp)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a        queue
		b        string
		expected string
		lenExp   int
	}{
		{queue{[]interface{}{"a", "b", "c", "d"}}, "e", "a", 5},
		{queue{[]interface{}{"a"}}, "b", "a", 2},
		{queue{[]interface{}{""}}, "", "", 2},
		{queue{[]interface{}{}}, "", "", 1},
	}
	for i, datest := range tests {
		datest.a.add(datest.b)
		actual := datest.a.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestPeek(t *testing.T) {
	var tests = []struct {
		a        queue
		expected string
	}{
		{queue{[]interface{}{"a", "b", "c", "d"}}, "a"},
		{queue{[]interface{}{"a"}}, "a"},
		{queue{[]interface{}{""}}, ""},
	}
	for i, datest := range tests {
		actual := datest.a.peek()
		actual2 := datest.a.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if actual != actual2 {
			t.Errorf("%v: queue altered from %v to %v", i, actual, actual2)
		}
	}
}
