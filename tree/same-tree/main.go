package main

import "fmt"

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.

// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

// Example 1.
// Input: p = [1,2,3], q = [1,2,3]
// Output: true

// Example 2.
// Input: p = [1,2], q = [1,null,2]
// Output: false

// Example 3.
// Input: p = [1,2,1], q = [1,1,2]
// Output: false

// Constraints:

// The number of nodes in both trees is in the range [0, 100].
// -104 <= Node.val <= 104

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

	n11 := &TreeNode{Val: 3}
	n22 := &TreeNode{Val: 9}
	n33 := &TreeNode{Val: 20}
	n11.Left = n22
	n11.Right = n33
	n44 := &TreeNode{Val: 15}
	n55 := &TreeNode{Val: 27}
	n22.Left = n44
	n22.Right = n55

	isSame := isSameTree(n1, n11)
	fmt.Println(isSame)
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
