package main

import "strconv"

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

func preOrder(root *TreeNode, s *string) {

	s1 := *s + strconv.Itoa(root.Val) + ","
	preOrder(root.Left, &s1)
	preOrder(root.Right, &s1)
}

func inOrder(root *TreeNode, s *string) {
	inOrder(root.Left, s)
	s1 := *s + strconv.Itoa(root.Val) + ","
	inOrder(root.Right, &s1)
}
