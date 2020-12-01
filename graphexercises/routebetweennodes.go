package graphexercises

import "hello/datastructureexercises"

//Given a directed graph, design an algorithm to find out whether there is a route between two nodes.

type node struct {
	data     interface{}
	children []*node
}

type directedGraph struct {
	nodes    []*node
	comparer datastructureexercises.Comparer
}

func newDirectedGraph(in map[interface{}][]interface{}, compare datastructureexercises.Comparer) *directedGraph {
	//m := make(map[interface{}]*node)
	//for k, v := range in {
	//}
	return nil
}

//Assume unique data; graph needs add, visited? (using a map)

func (g *directedGraph) findRoute(a, b interface{}) bool {
	return false
}
