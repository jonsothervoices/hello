package graphexercises

//Write a program to check if a binary tree is a binary search tree.

func (in bst) isValid() bool {
	if in.Left == nil && in.Right == nil {
		return true
	}
	if in.Left == nil {
		if in.Right.Data < in.Data {
			return false
		}
		return in.Right.isValid()
	}
	if in.Right == nil {
		if in.Left.Data > in.Data {
			return false
		}
		return in.Left.isValid()
	}
	if in.Right.Data < in.Data || in.Left.Data > in.Data {
		return false
	}
	return in.Right.isValid() && in.Left.isValid()
}
