package graphexercises

import (
	"hello/datastructureexercises"
	"testing"
)

func TestDepthFind(t *testing.T) {
	var tests = []struct {
		g        map[interface{}][]interface{}
		n        interface{}
		expected bool
	}{
		{map[interface{}][]interface{}{
			1: {11, 12, 13},
		}, 1, true},
		{map[interface{}][]interface{}{
			1: {11, 12, 13},
			2: {21, 22, 23},
			3: {31, 32, 33},
		}, 33, true},
		{map[interface{}][]interface{}{
			0: {1, 2, 3},
			1: {11, 12, 13},
			2: {21, 22, 23},
			3: {31, 32, 33},
		}, 0, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 12, 13},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 1, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 12, 13},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 43, false},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
			1:  {11, 12, 13},
		}, 1, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 2, 12, 13},
			2:  {1, 21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 31, true},
	}
	for i, datest := range tests {
		g := newDirectedGraph(datest.g, datastructureexercises.IntComparer)
		actual := g.depthFind(datest.n, nil)
		if actual == nil && datest.expected {
			t.Errorf("%v: actual %v, expected to find %v", i, actual, datest.n)
			t.FailNow()
		}
		if actual != nil && !datest.expected {
			t.Errorf("%v: actual %v was found, expected nothing", i, actual.data)
		}
		if actual != nil && datastructureexercises.IntComparer(actual.data, datest.n) != 0 {
			t.Errorf("%v: actual %v, expected %v", i, actual.data, datest.n)
		}
	}
}

func TestBreadthFind(t *testing.T) {
	var tests = []struct {
		g        map[interface{}][]interface{}
		n        interface{}
		expected bool
	}{
		{map[interface{}][]interface{}{
			1: {11, 12, 13},
		}, 1, true},
		{map[interface{}][]interface{}{
			1: {11, 12, 13},
			2: {21, 22, 23},
			3: {31, 32, 33},
		}, 33, true},
		{map[interface{}][]interface{}{
			0: {1, 2, 3},
			1: {11, 12, 13},
			2: {21, 22, 23},
			3: {31, 32, 33},
		}, 0, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 12, 13},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 1, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 12, 13},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 43, false},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			2:  {21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
			1:  {11, 12, 13},
		}, 1, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 2, 12, 13},
			2:  {1, 21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, 31, true},
		{map[interface{}][]interface{}{
			0:  {1, 2, 3},
			1:  {11, 2, 12, 13},
			2:  {1, 21, 22, 23},
			3:  {31, 32, 33},
			13: {0, 1, 2},
		}, nil, true},
	}
	for i, datest := range tests {
		g := newDirectedGraph(datest.g, datastructureexercises.IntComparer)
		actual := g.breadthFind(datest.n, nil)
		if actual == nil && datest.expected {
			t.Errorf("%v: actual %v, expected to find %v", i, actual, datest.n)
			t.FailNow()
		}
		if actual != nil && !datest.expected {
			t.Errorf("%v: actual %v was found, expected nothing", i, actual.data)
		}
		if actual != nil && datastructureexercises.IntComparer(actual.data, datest.n) != 0 {
			t.Errorf("%v: actual %v, expected %v", i, actual.data, datest.n)
		}
	}
}
