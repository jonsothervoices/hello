package graphexercises

//4.8: Write a program to find the first common ancestor in a binary tree. Avoid storing additional nodes in a data structure (NOTE: this is not necessarily a binary search tree).

//commonAncestor returns -1 if no ancestor exists
//commonAncestor has time O(nlogn)
func (t *bst) commonAncestor(a, b int) int {
	//declare aNode and bNode
	aNode := t.tFind(a)
	//find a and set it to aNode
	for aNode != nil {
		if f := aNode.tFind(b); f != nil {
			return aNode.Data
		}
		aNode = aNode.parent
	}
	return -1
}

//tFind finds a node in any binary tree, not just bst.
//tFind has time O(n)
func (t *bst) tFind(i int) *bst {
	if t.Data == i {
		return t
	}
	if t.Left != nil {
		l := t.Left.find(i)
		//conditional return because we still need to check the right
		if l != nil {
			return l
		}
	}
	if t.Right != nil {
		return t.Right.find(i)
	}
	//if we haven't found it by now, we're at a dead end
	return nil
}
