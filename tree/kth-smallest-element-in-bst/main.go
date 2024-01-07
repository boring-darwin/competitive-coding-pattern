package main

import "fmt"

// Given the root of a binary search tree, and an integer k,
// return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

// Example 1:
// Input: root = [3,1,4,null,2], k = 1
// Output: 1

// Example 2:
// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3

// Constraints:

// The number of nodes in the tree is n.
// 1 <= k <= n <= 104
// 0 <= Node.val <= 104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 8}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 14}
	n1.Left = n2
	n1.Right = n3
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 7}
	n2.Left = n4
	n2.Right = n5
	n6 := &TreeNode{Val: 11}
	n7 := &TreeNode{Val: 28}
	n3.Left = n6
	n3.Right = n7

	// isSubTree := findKthSmallestElement(n1)
	arr := make([]int, 0)
	createSortedArrayFromBST(n1, &arr)
	ans := findKthSmallestElement(1, arr)
	fmt.Printf("\n %d", ans)
}

// Inorder Traversal of tree will generate sorted array
func createSortedArrayFromBST(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	createSortedArrayFromBST(root.Left, arr)
	fmt.Printf(" %d", root.Val)
	*arr = append(*arr, root.Val)
	createSortedArrayFromBST(root.Right, arr)
}

func findKthSmallestElement(k int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if i+1 == k {
			return arr[i]
		}
	}

	return 0
}
