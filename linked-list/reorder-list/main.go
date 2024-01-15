package main

import "fmt"

// You are given the head of a singly linked-list. The list can be represented as:

// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {
	// n5 := &ListNode{Val: 6, Next: nil}
	// n4 := &ListNode{Val: 5, Next: nil}
	// n3 := &ListNode{Val: 4, Next: n4}
	// n2 := &ListNode{Val: 3, Next: n3}
	// n1 := &ListNode{Val: 2, Next: nil}
	head1 := &ListNode{Val: 1, Next: nil}

	middleHead := getMiddleOfLinkedList(head1)
	print(head1)
	reverseHead := reverLinkedList(middleHead)
	reorder(head1, reverseHead)
	fmt.Println()
	print(head1)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func reorder(head, reverse *ListNode) {

	for reverse.Next != nil {
		nextNode := head.Next
		reverseNextNode := reverse.Next
		head.Next = reverse
		reverse.Next = nextNode
		head = nextNode
		reverse = reverseNextNode
	}
}

func getMiddleOfLinkedList(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil {
		if fast.Next == nil {
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head.Next
	head.Next = nil
	prev := head
	var temp *ListNode
	for curr != nil {
		temp = curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
