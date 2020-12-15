package graphexercises

//given a sorted array with unique integer elements, write a program to create a binary search tree with miunimal height.

//Binary search tree: no more than two children per node, all Left children <=n,all Right children, where n is the value of the node element.

type bst struct {
	Data   int
	Left   *bst
	Right  *bst
	parent *bst
}

func newBST(in []int, parent *bst) *bst {
	if len(in) == 0 {
		return nil
	}
	if len(in) == 1 {
		return &bst{Data: in[0], parent: parent}
	}
	mid := len(in) / 2
	next := &bst{Data: in[mid], parent: parent}
	next.Left = newBST(in[0:mid], next)
	next.Right = newBST(in[mid+1:], next)
	return next
}
