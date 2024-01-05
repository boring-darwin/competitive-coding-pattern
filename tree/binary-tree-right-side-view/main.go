package main

import "fmt"

// https://leetcode.com/problems/binary-tree-right-side-view/description/
// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.

// Example 1:
// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]

// Example 2:
// Input: root = [1,null,3]
// Output: [1,3]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3
	n4 := &TreeNode{Val: 4}
	n2.Left = n4

	out := rightSideView(n1)
	fmt.Print(out)
}

func rightSideView(node *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	output := make([]int, 0)
	queue = append(queue, node)
	if node == nil {
		return output
	}
	output = append(output, node.Val)
	for len(queue) > 0 {
		currLen := len(queue)
		done := false
		for i := 0; i < currLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if (node.Right != nil && !done) || (node.Left != nil && !done) {
				if node.Right != nil {
					output = append(output, node.Right.Val)
				} else if node.Left != nil {
					output = append(output, node.Left.Val)
				}
				done = true
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return output
}
