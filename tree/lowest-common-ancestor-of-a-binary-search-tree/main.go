package main

import "fmt"

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest
// node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

// Example 1:
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// Output: 6
// Explanation: The LCA of nodes 2 and 8 is 6.

// Example 2:
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// Output: 2
// Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

// Example 3:
// Input: root = [2,1], p = 2, q = 1
// Output: 2

// Constraints:

// The number of nodes in the tree is in the range [2, 105].
// -109 <= Node.val <= 109
// All Node.val are unique.
// p != q
// p and q will exist in the BST.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 8}
	n1.Left = n2
	n1.Right = n3
	n4 := &TreeNode{Val: 0}
	n5 := &TreeNode{Val: 4}
	n2.Left = n4
	n2.Right = n5
	n6 := &TreeNode{Val: 7}
	n7 := &TreeNode{Val: 9}
	n3.Left = n6
	n3.Right = n7
	n8 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 5}
	n5.Left = n8
	n5.Right = n9

	printTree(n1)
	out := lowestCommonAncestorOfABinarySearchTree(n1, n2, n5)
	fmt.Println(out.Val)
}

func lowestCommonAncestorOfABinarySearchTree(root *TreeNode, p, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if curr.Val > p.Val && curr.Val > q.Val {
			curr = curr.Left
		} else if curr.Val < p.Val && curr.Val < q.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}
	return curr
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf(" %d", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}
