package graphexercises

//Write a program to find the "next" node (ie in-order successor) of a given node of a BST.

func (b *bst) successor() *bst {
	if b == nil {
		return nil
	}
	if b.Right != nil {
		current := b.Right
		for current.Left != nil {
			current = current.Left
		}
		return current
	}
	if b.parent != nil {
		if b.parent.Data > b.Data {
			return b.parent
		}
		current := b.parent
		for current.parent != nil {
			if current.parent.Data > b.Data {
				return current.parent
			}
			current = current.parent
		}
	}
	return nil
}

func (b *bst) find(i int) *bst {
	if b == nil {
		return nil
	}
	if b.Data == i {
		return b
	}
	if b.Data < i && b.Right != nil {
		return b.Right.find(i)
	}
	if b.Data > i && b.Left != nil {
		return b.Left.find(i)
	}
	return nil
}
