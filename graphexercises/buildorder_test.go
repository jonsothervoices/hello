package graphexercises

import (
	// "fmt"
	"hello/datastructureexercises"
	"testing"
)

func TestBuildOrder(t *testing.T) {
	var tests = []struct {
		g        map[interface{}][]interface{}
		isValid  bool
		expected []interface{}
	}{
		{map[interface{}][]interface{}{
			"e": {},
			"c": {"d"},
			"a": {"f"},
			"b": {"f"},
			"d": {"b", "a"},
			"f": {}}, true,
			[]interface{}{"f", "a", "b", "d", "c", "e"}},
		{map[interface{}][]interface{}{
			"e": {},
			"c": {"d"},
			"a": {"f"},
			"b": {"f"},
			"d": {"b", "a"},
			"f": {"d"}}, false,
			nil},
	}
	for i, datest := range tests {
		g := newDirectedGraph(datest.g, datastructureexercises.StrComparer)
		// fmt.Printf("%v\n", g)
		actual, err := g.buildOrder()
		if err == nil && !datest.isValid {
			t.Errorf("%v: expected no valid order", i)
		}
		if err != nil && datest.isValid {
			t.Errorf("%v: error: %v: expected valid order", i, err)
		}
		if len(datest.expected) != len(actual) {
			t.Errorf("%v: actual length %v, expected length %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range datest.expected {
			if v != actual[j] {
				t.Errorf("%v: actual[%v] %v, expected %v", i, j, actual[j], v)
			}
		}
	}
}
