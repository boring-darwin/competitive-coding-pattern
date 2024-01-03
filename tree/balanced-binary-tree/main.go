package main

import "fmt"

// Given a binary tree, determine if it is height-balanced
// (A height-balanced binary tree is a binary tree in which the depth of the two subtrees
// of every node never differs by more than one.)

// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: true

// Example 2:
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false

// Example 3:
// Input: root = []
// Output: true

// Constraints:
// The number of nodes in the tree is in the range [0, 5000].
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

	_, isBalanced := depth(n1)
	fmt.Println(isBalanced)
}

func depth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	d1, d1Bal := depth(root.Left)
	d2, d2Bal := depth(root.Right)
	isBalanced := true
	if (d1-d2 == -1 || d2-d1 == -1) || (d1-d2 == 0) {
		isBalanced = true
	} else {
		isBalanced = false
	}

	return max(d1, d2) + 1, d1Bal && d2Bal && isBalanced
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
