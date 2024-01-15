package main

import "fmt"

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// Example 2:
// Input: list1 = [], list2 = []
// Output: []

// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]

// Constraints:
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {
	n4 := &ListNode{Val: 98, Next: nil}
	n3 := &ListNode{Val: 79, Next: n4}
	n2 := &ListNode{Val: 68, Next: n3}
	n1 := &ListNode{Val: 39, Next: n2}
	head1 := &ListNode{Val: 10, Next: n1}

	n44 := &ListNode{Val: 106, Next: nil}
	n33 := &ListNode{Val: 57, Next: n44}
	n23 := &ListNode{Val: 48, Next: n33}
	n11 := &ListNode{Val: 29, Next: n23}
	head2 := &ListNode{Val: 10, Next: n11}

	mergedListHead := mergeSortedList(head1, head2)
	print(mergedListHead)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func mergeSortedList(head1, head2 *ListNode) *ListNode {
	var mergedListHead *ListNode
	var prev *ListNode
	node1 := head1
	node2 := head2
	for node1 != nil && node2 != nil {
		newMergedListNode := &ListNode{}
		if node1.Val < node2.Val {
			newMergedListNode.Val = node1.Val
			node1 = node1.Next
		} else {
			newMergedListNode.Val = node2.Val
			node2 = node2.Next
		}
		if mergedListHead == nil {
			mergedListHead = newMergedListNode
			prev = newMergedListNode
			continue
		}
		prev.Next = newMergedListNode
		prev = newMergedListNode
	}

	for node1 != nil {
		newMergedListNode := &ListNode{Val: node1.Val}
		prev.Next = newMergedListNode
		prev = newMergedListNode
		node1 = node1.Next
	}

	for node2 != nil {
		newMergedListNode := &ListNode{Val: node2.Val}
		prev.Next = newMergedListNode
		prev = newMergedListNode
		node2 = node2.Next
	}

	return mergedListHead
}
