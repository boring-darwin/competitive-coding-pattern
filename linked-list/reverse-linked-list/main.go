package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

func main() {
	head := &Node{Val: 10}
	n1 := &Node{Val: 9}
	n2 := &Node{Val: 8}
	n3 := &Node{Val: 7}
	n4 := &Node{Val: 6}

	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	print(head)
	newHead := reverse(head)
	fmt.Println()
	print(newHead)
}

func print(head *Node) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func reverse(head *Node) *Node {
	if head == nil {
		return head
	}
	
	curr := head.Next
	head.Next = nil
	prev := head
	var temp *Node
	for curr != nil {
		temp = curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
