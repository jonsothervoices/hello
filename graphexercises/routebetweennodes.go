package graphexercises

import (
	"fmt"
	"hello/datastructureexercises"
	"sort"
)

//Given a directed graph, design an algorithm to find out whether there is a route between two nodes.

type node struct {
	data     interface{}
	children []*node
	comparer datastructureexercises.Comparer
}

func newDirectedGraph(in map[interface{}][]interface{}, compare datastructureexercises.Comparer) *node {
	m := make(map[interface{}]*node)
	g := &node{comparer: compare}
	s := []interface{}{}
	for k := range in {
		s = append(s, k)
	}
	sort.Slice(s, func(i, j int) bool {
		s1 := fmt.Sprintf("%v", s[i])
		s2 := fmt.Sprintf("%v", s[j])
		return s1 < s2
	})
	for _, k := range s {
		parent, ok := m[k]
		if !ok {
			parent = &node{data: k, comparer: compare}
			m[k] = parent
		}
		for _, c := range in[k] {
			child, ok := m[c]
			if !ok {
				child = &node{data: c, comparer: compare}
				m[c] = child
			}
			parent.children = append(parent.children, child)
		}
		g.children = append(g.children, parent)
	}
	return g
}

//Assume unique data; graph needs add, visited? (using a map)
func (g *node) findNode(a interface{}, check map[interface{}]bool) *node {
	if check == nil {
		check = make(map[interface{}]bool)
	}
	if g.data != nil && g.comparer(a, g.data) == 0 {
		return g
	}
	for _, v := range g.children {
		if check[v.data] {
			continue
		}
		check[v.data] = true
		f := v.findNode(a, check)
		if f != nil {
			return f
		}
	}
	return nil
}

func (g *node) findRoute(a, b interface{}) (bool, error) {
	if g.comparer(a, b) == 0 {
		return true, nil
	}
	na := g.findNode(a, nil)
	if na == nil {
		return false, fmt.Errorf("%v is not in graph", a)
	}
	nb := na.findNode(b, nil)
	return nb != nil, nil
}
