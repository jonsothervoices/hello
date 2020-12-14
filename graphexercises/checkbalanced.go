package graphexercises

import "math"

//Write a program to check if a binary tree is balanced (ie. the heights of 2 subtrees of any node never differ by more than 1).
func (in bst) isBalanced() (bool, int) {
	//if depth>1{
	//return false
	//}
	if in.Left == nil {
		if in.Right == nil {
			return true, 1
		}
		return in.Right.Left == nil && in.Right.Right == nil, 2
	}
	if in.Right == nil {
		if in.Left == nil {
			return true, 1
		}
		return in.Left.Left == nil && in.Left.Right == nil, 2
	}
	l, ldepth := in.Left.isBalanced()
	r, rdepth := in.Right.isBalanced()
	b := l && r && math.Abs(float64(ldepth-rdepth)) < 2
	d := int(math.Max(float64(ldepth), float64(rdepth))) + 1
	return b, d
}
