package main

import "fmt"

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along
// the longest path from the root node down to the farthest leaf node.

// Example: 1
// Input: root = [3,9,20,null,null,15,7]
// Output: 3

// Example 2:
// Input: root = [1,null,2]
// Output: 2

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

	deep := depth(n1)
	fmt.Println(deep)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(depth(node.Left), depth(node.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
