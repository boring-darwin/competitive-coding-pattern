package main

import "fmt"

// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
// The path sum of a path is the sum of the node's values in the path.
// Given the root of a binary tree, return the maximum path sum of any non-empty path.

// Example 1:
// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

// Example 2:
// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

// Constraints:
// The number of nodes in the tree is in the range [1, 3 * 104].
// -1000 <= Node.val <= 1000

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

	var dfs func(root *TreeNode) int

	res := n1.Val
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		leftMax := max(left, 0)
		rightMax := max(right, 0)

		res = max(res, root.Val+leftMax+rightMax)
		return root.Val + max(leftMax, rightMax)
	}
	dfs(n1)

	fmt.Println(res)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
