package main

import "fmt"

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

// Example 2:
// Input: head = [1], n = 1
// Output: []

// Example 3:
// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {
	// n4 := &ListNode{Val: 98, Next: nil}
	// n3 := &ListNode{Val: 79, Next: n4}
	// n2 := &ListNode{Val: 68, Next: n3}
	n1 := &ListNode{Val: 39, Next: nil}
	head1 := &ListNode{Val: 10, Next: n1}

	print(head1)
	removeNthNodeFromEnd(head1, 2)
	fmt.Println()
	print(head1)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func removeNthNodeFromEnd(head *ListNode, n int) {

	if head.Next == nil && n == 1 {
		head = &ListNode{}
		return
	}
	beforeEnd := head
	end := head

	for i := 1; i <= n; i++ {
		end = end.Next
	}

	prev := beforeEnd
	// if we have only 2 nodes.
	if end == nil {
		head = head.Next
		return
	}

	for end != nil {
		prev = beforeEnd
		beforeEnd = beforeEnd.Next
		end = end.Next
	}

	prev.Next = beforeEnd.Next
}
