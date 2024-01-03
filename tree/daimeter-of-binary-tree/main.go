package main

import "fmt"

// Given the root of a binary tree, return the length of the diameter of the tree.

// The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
// This path may or may not pass through the root.

// The length of a path between two nodes is represented by the number of edges between them.

// Example 1:
// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

// Example 2:
// Input: root = [1,2]
// Output: 1

// Constraints:
// The number of nodes in the tree is in the range [1, 104].
// -100 <= Node.val <= 100

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

	height := 0
	deep := width(n1, &height)
	fmt.Println(deep)
}

func width(node *TreeNode, height *int) int {
	if node == nil {
		return 0
	}

	left := width(node.Left, height)
	right := width(node.Right, height)
	*height = max(*height, left+right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
