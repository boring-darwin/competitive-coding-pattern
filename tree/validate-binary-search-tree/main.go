package main

import "fmt"

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).

// A valid BST is defined as follows:

// The left
// subtree
//  of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

// Example 1:
// Input: root = [2,1,3]
// Output: true

// Example 2:
// Input: root = [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.

// Constraints:
// The number of nodes in the tree is in the range [1, 104].
// -231 <= Node.val <= 231 - 1

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

	isSubTree := isValidBST(n1)
	fmt.Println(isSubTree)
}

func isValidBST(root *TreeNode) bool {
	min := -2147483648
	max := 2147483647

	return validate(root, max, min)
}

func validate(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val < min || root.Val > max {
		return false
	}

	return validate(root.Left, root.Val-1, min) && validate(root.Right, max, root.Val+1)
}
