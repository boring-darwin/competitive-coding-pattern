package main

import "fmt"

// Given the root of a binary tree, invert the tree, and return its root.

// Example 1:
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]

// Example 2:
// Input: root = [2,1,3]
// Output: [2,3,1]

// Example 3:
// Input: root = []
// Output: []

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {

	n1 := &Node{data: 4}
	n2 := &Node{data: 2}
	n3 := &Node{data: 7}
	n1.left = n2
	n1.right = n3
	n4 := &Node{data: 1}
	n5 := &Node{data: 3}
	n2.left = n4
	n2.right = n5
	n6 := &Node{data: 6}
	n7 := &Node{data: 9}
	n3.left = n6
	n3.right = n7

	printTree(n1)
	fmt.Println()
	invertNode(n1)
	printTree(n1)
}

func invertNode(node *Node) {
	if node == nil {
		return
	}

	temp := node.left
	node.left = node.right
	node.right = temp

	invertNode(node.left)
	invertNode(node.right)
}

func printTree(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf(" %d", node.data)
	printTree(node.left)
	printTree(node.right)
}
