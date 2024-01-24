package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}

	root := constructTree(inOrder, preOrder)
	printTree(root)

}

func constructTree(inorder, preorder []int) *TreeNode {

	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	newNode := &TreeNode{
		Val: preorder[0],
	}

	inIndex := search(inorder, newNode.Val)

	newNode.Left = constructTree(inorder[:inIndex], preorder[1:inIndex+1])
	newNode.Right = constructTree(inorder[inIndex+1:], preorder[inIndex+1:])

	return newNode
}

func search(inorder []int, num int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == num {
			return i
		}
	}
	return -1
}

func printTree(node *TreeNode) {
	if node == nil {
		return
	}

	printTree(node.Left)
	fmt.Printf(" %d", node.Val)
	printTree(node.Right)
}
