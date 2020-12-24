package graphexercises

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type bst struct {
	Data   int
	Left   *bst
	Right  *bst
	parent *bst
}

//4.2: given a sorted array with unique integer elements, write a program to create a binary search tree with miunimal height.
//newBST constructor: will return a binary search tree only if the input slice is sorted; otherwise it will just be a binary tree.
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

//4.4: Write a program to check if a binary tree is balanced (ie. the heights of 2 subtrees of any node never differ by more than 1).
func (t bst) isBalanced() (bool, int) {
	if t.Left == nil {
		if t.Right == nil {
			return true, 1
		}
		return t.Right.Left == nil && t.Right.Right == nil, 2
	}
	if t.Right == nil {
		if t.Left == nil {
			return true, 1
		}
		return t.Left.Left == nil && t.Left.Right == nil, 2
	}
	l, ldepth := t.Left.isBalanced()
	r, rdepth := t.Right.isBalanced()
	b := l && r && math.Abs(float64(ldepth-rdepth)) < 2
	d := int(math.Max(float64(ldepth), float64(rdepth))) + 1
	return b, d
}

//4.5: Write a program to check if a binary tree is a binary search tree.
func (t bst) isValid() bool {
	if t.Left == nil && t.Right == nil {
		return true
	}
	if t.Left == nil {
		if t.Right.Data < t.Data {
			return false
		}
		return t.Right.isValid()
	}
	if t.Right == nil {
		if t.Left.Data > t.Data {
			return false
		}
		return t.Left.isValid()
	}
	if t.Right.Data < t.Data || t.Left.Data > t.Data {
		return false
	}
	return t.Right.isValid() && t.Left.isValid()
}

//4.6: Write a program to find the "next" node (ie in-order successor) of a given node of a BST.
func (t *bst) successor() *bst {
	if t == nil {
		return nil
	}
	if t.Right != nil {
		current := t.Right
		for current.Left != nil {
			current = current.Left
		}
		return current
	}
	if t.parent != nil {
		if t.parent.Data > t.Data {
			return t.parent
		}
		current := t.parent
		for current.parent != nil {
			if current.parent.Data > t.Data {
				return current.parent
			}
			current = current.parent
		}
	}
	return nil
}

func (t *bst) find(n int) *bst {
	if t == nil {
		return nil
	}
	if t.Data == n {
		return t
	}
	if t.Right != nil && t.Data < n {
		return t.Right.find(n)
	}
	if t.Left != nil {
		return t.Left.find(n)
	}
	return nil
}

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
	if t == nil || t.Data == i {
		return t
	}
	if t.Left != nil {
		l := t.Left.tFind(i)
		//conditional return because we still need to check the right
		if l != nil {
			return l
		}
	}
	if t.Right != nil {
		return t.Right.tFind(i)
	}
	//if we haven't found it by now, we're at a dead end
	return nil
}

//4.11 Random Node: You are implementing a binary tree class from scratch which, in addition to "insert", "find", and "delete", has a method "getRandomNode()" which returns a random node from the tree. All nodes should be equally likely to be chosen. Design and implement an algorithm for "getRandomNode", and explain how you would implement the rest of the methods.

//getRandomNode has time O(n)
func (t *bst) getRandomNode() (int, error) {
	//write elements of t to a slice
	s := t.traverse()
	// zero check
	if len(s) == 0 {
		return -1, fmt.Errorf("Empty tree!")
	}
	//Seed RNG
	rand.Seed(time.Now().UnixNano())
	// call rand bounded by slice length
	rInt := intn(len(s))
	// fmt.Println(rInt, s)
	//return slice value at index from rand
	return s[rInt], nil
}

//binary tree traversal--returns a slice containing all elements of t.
func (t *bst) traverse() (out []int) {
	//nil check
	if t == nil {
		return
	}
	if t.Left != nil {
		out = append(out, t.Left.traverse()...)
	}
	out = append(out, t.Data)
	if t.Right != nil {
		out = append(out, t.Right.traverse()...)
	}
	return out
}

//stubbing rand
var intn = rand.Intn

//4.12 paths with sum: Given a binary search tree with intgewr elements, write a program to count the nuymber of paths that sum to a given value. The path does not need to start or end at the root or leaf, but it must go downwards.

//t=4, 2, 1
//n=10, 6, 4
//ret=0, 0, 0

func (t *bst) pathSum(n int) int {
	return t.doPathSum(n, n)
}

func (t *bst) doPathSum(n, c int) int {
	// fmt.Printf("starting from node %v, looking for %v\n", t.Data, c)
	ret := 0
	if t == nil {
		return 0
	}
	if t.Data == n {
		ret++
		fmt.Printf("found a path at node %v\n", t.Data)
	}
	if t.Left != nil {
		// if t.Data+t.Left.Data == n {
		// 	ret++
		// 	fmt.Printf("found a path from node %v to node %v\n", t.Data, t.Left.Data)
		// }
		fmt.Printf("Checking %v for altered sum %v\n", t.Left.Data, n-t.Data)
		ret += t.Left.doPathSum(n-t.Data, c)
		fmt.Printf("Checking %v for original sum %v\n", t.Left.Data, c)
		ret += t.Left.doPathSum(c, c)
	}
	if t.Right != nil && t.Data <= n {
		// if t.Data+t.Right.Data == n {
		// 	fmt.Printf("found a path from node %v to node %v\n", t.Data, t.Right.Data)
		// 	ret++
		// }
		fmt.Printf("Checking %v for altered sum %v\n", t.Right.Data, n-t.Data)
		ret += t.Right.doPathSum(n-t.Data, c)
		fmt.Printf("Checking %v for original sum %v\n", t.Right.Data, c)
		ret += t.Right.doPathSum(c, c)
	}
	return ret
}
