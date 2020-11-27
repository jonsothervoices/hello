package datastructureexercises

import "testing"

func TestRemove(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		lenExp   int
		expected interface{}
		qtype    string
	}{
		{[]interface{}{"a", "b", "c", "d", "e"}, 4, "a", simpleQueueType},
		{[]interface{}{"a"}, 0, "a", simpleQueueType},
		{[]interface{}{""}, 0, "", simpleQueueType},
		{[]interface{}{}, 0, nil, simpleQueueType},
		{[]interface{}{"a", "b", "c", "d", "e"}, 4, "a", stackQueueType},
		{[]interface{}{"a"}, 0, "a", stackQueueType},
		{[]interface{}{""}, 0, "", stackQueueType},
		{[]interface{}{}, 0, nil, stackQueueType},
	}
	for i, datest := range tests {
		q := queueFactory(datest.qtype, datest.a)
		actual := q.remove()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if q.len() != datest.lenExp {
			t.Errorf("%v: length %v, expected %v", i, q.len(), datest.lenExp)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		b        string
		expected string
		lenExp   int
		qtype    string
	}{
		{[]interface{}{"a", "b", "c", "d"}, "e", "a", 5, simpleQueueType},
		{[]interface{}{"a"}, "b", "a", 2, simpleQueueType},
		{[]interface{}{""}, "", "", 2, simpleQueueType},
		{[]interface{}{}, "", "", 1, simpleQueueType},
		{[]interface{}{"a", "b", "c", "d"}, "e", "a", 5, stackQueueType},
		{[]interface{}{"a"}, "b", "a", 2, stackQueueType},
		{[]interface{}{""}, "", "", 2, stackQueueType},
		{[]interface{}{}, "", "", 1, stackQueueType},
	}
	for i, datest := range tests {
		q := queueFactory(simpleQueueType, datest.a)
		q.add(datest.b)
		actual := q.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestPeek(t *testing.T) {
	var tests = []struct {
		a        []interface{}
		expected interface{}
		qtype    string
	}{
		{[]interface{}{"a", "b", "c", "d"}, "a", simpleQueueType},
		{[]interface{}{"a"}, "a", simpleQueueType},
		{[]interface{}{""}, "", simpleQueueType},
		{[]interface{}{"a", "b", "c", "d"}, "a", stackQueueType},
		{[]interface{}{"a"}, "a", stackQueueType},
		{[]interface{}{""}, "", stackQueueType},
	}
	for i, datest := range tests {
		q := queueFactory(simpleQueueType, datest.a)
		actual := q.peek()
		actual2 := q.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if actual != actual2 {
			t.Errorf("%v: queue altered from %v to %v", i, actual, actual2)
		}
	}
}
