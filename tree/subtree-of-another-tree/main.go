package main

import "fmt"

// Given the roots of two binary trees root and subRoot,
// return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
// The tree tree could also be considered as a subtree of itself.

// Example 1:
// Input: root = [3,4,5,1,2], subRoot = [4,1,2]
// Output: true

// Example 2:
// Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
// Output: false

// Constraints:

// The number of nodes in the root tree is in the range [1, 2000].
// The number of nodes in the subRoot tree is in the range [1, 1000].
// -104 <= root.val <= 104
// -104 <= subRoot.val <= 104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 20}
	n1.Left = n2
	n1.Right = n3
	n4 := &TreeNode{Val: 15}
	n5 := &TreeNode{Val: 27}
	n2.Left = n4
	n2.Right = n5

	isSubTree := subTree(n1, n2)
	fmt.Println(isSubTree)
}

func subTree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || isSameTree(root.Left, subRoot) || isSameTree(root.Right, subRoot)
}

func isSameTree(root1, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}

	return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
}
