package graphexercises

import (
	"fmt"
	"hello/datastructureexercises"
	"sort"
)

type node struct {
	data     interface{}
	children []*node
	comparer datastructureexercises.Comparer
}

//Directed graph constructor
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

//Finds a node within a graph(*node), assuming unique data.
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

//4.1 Given a directed graph, design an algorithm to find out whether there is a route between two nodes.
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

//4.7: You are given a list of projects and a list of dependencies (which is a list of pairs of  projects, where the first project is dependent on the second project). All of a project's dependencies must be built before the project is. Find a build order that will allow the projects to be built. If there is no valid build order, return an error.
// Example:
//Input:
// Projects: a,b,c,d,e,f
// dependencies: (d,a), (b,f), (d,b), (a,f), (c,d)
// Output: f,e,a,b,d,c

func (g *node) buildOrder() ([]interface{}, error) {
	m := make(map[interface{}]int)
	err := g.doBuildOrder(m)
	if err != nil {
		return nil, err
	}
	ret := make([]interface{}, len(m))
	for k, v := range m {
		ret[v] = k
	}
	return ret, nil
}

func (g *node) doBuildOrder(in map[interface{}]int) error {
	//if we've already seen it, skip
	_, ok := in[g.data]
	if ok {
		return nil
	}
	//iterate over children
	for _, v := range g.children {
		//if we can find g in v, return error
		// fmt.Printf("ranging over %v, checking %v\n", g.data, v.data)
		c := v.findNode(g.data, nil)
		if c != nil {
			return fmt.Errorf("No valid build order exists: '%v' and '%v' have circular dependency", g.data, v.data)
		}
		//no loop, recurse
		err := v.doBuildOrder(in)
		if err != nil {
			return err
		}
	}
	//add data to output map
	if g.data != nil {
		// fmt.Printf("adding %v to map at location %v\n", g.data, len(in))
		in[g.data] = len(in)
	}
	//return  no error
	return nil
}
