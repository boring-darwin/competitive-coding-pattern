package main

import "fmt"

// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).

// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []

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
	n5 := &TreeNode{Val: 7}
	n3.Left = n4
	n3.Right = n5

	output := levelOrderTraversal(n1)
	fmt.Println(output)
}

func levelOrderTraversal(node *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	output := make([][]int, 0)
	queue = append(queue, node)
	if node == nil {
		return output
	}
	o := make([]int, 0)
	o = append(o, node.Val)
	output = append(output, o)
	for len(queue) > 0 {
		currLen := len(queue)
		arr := make([]int, 0)
		for i := 0; i < currLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				arr = append(arr, node.Left.Val)
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				arr = append(arr, node.Right.Val)
				queue = append(queue, node.Right)
			}
		}
		if len(arr) > 0 {
			output = append(output, arr)
		}
	}
	return output
}
