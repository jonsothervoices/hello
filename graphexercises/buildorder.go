package graphexercises

import "fmt"

// You are given a list of projects and a list of dependencies (which is a list of pairs of  projects, where the first project is dependent on the second project). All of a project's dependencies must be built before the project is. Find a build order that will allow the projects to be built. If there is no valid build order, return an error.
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
		fmt.Printf("ranging over %v, checking %v\n", g.data, v.data)
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
		fmt.Printf("adding %v to map at location %v\n", g.data, len(in))
		in[g.data] = len(in)
	}
	//return  no error
	return nil
}
