package graphexercises

import (
	"fmt"
	"hello/datastructureexercises"
	"sort"
)

// You are given a list of projects and a list of dependencies (which is a list of pairs of  projects, where the first project is dependent on the second project). All of a project's dependencies must be built before the project is. Find a build order that will allow the projects to be built. If there is no valid build order, return an error.
// Example:
//Input:
// Projects: a,b,c,d,e,f
// dependencies: (d,a), (b,f), (d,b), (a,f), (c,d)
// Output: f,e,a,b,d,c

// func (g *node) String() string {
// 	c := []string{}
// 	for _, v := range g.children {
// 		c = append(c, v.String())
// 	}
// 	return fmt.Sprintf(`{"Data": "%v", "children": [%v]}`, g.data, strings.Join(c, ","))
// }

func newOrderGraph(in map[interface{}][]interface{}, compare datastructureexercises.Comparer) *node {
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

func (g *node) buildOrder() ([]interface{}, error) {
	m := make(map[interface{}]int)
	err := g.orderTrav(m)
	if err != nil {
		return nil, err
	}
	out := make([]interface{}, len(m))
	for k, v := range m {
		out[v] = k
	}
	return out, nil
}

func (g *node) orderTrav(in map[interface{}]int) error {
	if _, ok := in[g.data]; ok {
		return nil
	}
	for _, v := range g.children {
		if v.findNode(g.data, nil) != nil {
			return fmt.Errorf("no valid order: %v and %v are codependent", g.data, v.data)
		}
		err := v.orderTrav(in)
		if err != nil {
			return err
		}
	}
	if g.data != nil {
		in[g.data] = len(in)
	}
	return nil
}
