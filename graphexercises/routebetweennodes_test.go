package graphexercises

import (
	"hello/datastructureexercises"
	"testing"
)

func TestFindRoute(t *testing.T) {
	var tests = []struct {
		g        map[interface{}][]interface{}
		a        interface{}
		b        interface{}
		expected bool
	}{
		{map[interface{}][]interface{}{
			1: {11, 12, 13},
		}, 1, 13, true},
		{map[interface{}][]interface{}{
			1: {11, 12, 13},
			2: {21, 22, 23},
			3: {31, 32, 33},
		}, 1, 33, false},
		{map[interface{}][]interface{}{
			0: {1, 2, 3},
			1: {11, 12, 13},
			2: {21, 22, 23},
			3: {31, 32, 33},
		}, 1, 33, true},
		{map[interface{}][]interface{}{
			0:  {1, 2},
			1:  {11, 12, 13},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 1, 33, true},
	}
	for i, datest := range tests {
		g := newDirectedGraph(datest.g, datastructureexercises.IntComparer)
		actual := g.findRoute(datest.a, datest.b)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
