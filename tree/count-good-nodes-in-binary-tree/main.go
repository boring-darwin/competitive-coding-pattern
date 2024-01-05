package main

import "fmt"

// Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.

// Example 1:
// Input: root = [3,1,4,3,null,1,5]
// Output: 4
// Explanation: Nodes in blue are good.
// Root Node (3) is always a good node.
// Node 4 -> (3,4) is the maximum value in the path starting from the root.
// Node 5 -> (3,4,5) is the maximum value in the path
// Node 3 -> (3,1,3) is the maximum value in the path.

// Example 2:
// Input: root = [3,3,null,4,2]
// Output: 3
// Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

// Example 3:
// Input: root = [1]
// Output: 1
// Explanation: Root is considered as good.

// Constraints:
// The number of nodes in the binary tree is in the range [1, 10^5].
// Each node's value is between [-10^4, 10^4].

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3
	n4 := &TreeNode{Val: 8}
	n2.Left = n4

	ans := countGoodNodes(n1, n1.Val)
	fmt.Print(ans)
}

func countGoodNodes(node *TreeNode, recentBig int) int {
	if node == nil {
		return 0
	}

	res := 1
	max := node.Val

	if recentBig > node.Val {
		res = 0
		max = recentBig
	}

	res += countGoodNodes(node.Left, max)
	res += countGoodNodes(node.Right, max)

	return res
}
