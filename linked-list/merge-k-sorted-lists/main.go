package main

import "fmt"

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.

// Example 1:
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

// Example 2:
// Input: lists = []
// Output: []

// Example 3:
// Input: lists = [[]]
// Output: []

// Constraints:
// k == lists.length
// 0 <= k <= 104
// 0 <= lists[i].length <= 500
// -104 <= lists[i][j] <= 104
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 104.

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

	list := []*ListNode{head1, head2}
	mergedListHead := mergeKSortedList(list)
	print(mergedListHead)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func mergeKSortedList(list []*ListNode) *ListNode {
	if len(list) == 0 {
		return nil
	}
	var prevList *ListNode = list[0]
	for i := 1; i < len(list); i++ {
		var mergedListHead *ListNode
		var prev *ListNode
		node1 := list[i]
		node2 := prevList
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
			if mergedListHead == nil {
				mergedListHead = node1
				break
			}
			newMergedListNode := &ListNode{Val: node1.Val}
			prev.Next = newMergedListNode
			prev = newMergedListNode
			node1 = node1.Next
		}

		for node2 != nil {
			if mergedListHead == nil {
				mergedListHead = node2
				break
			}
			newMergedListNode := &ListNode{Val: node2.Val}
			prev.Next = newMergedListNode
			prev = newMergedListNode
			node2 = node2.Next
		}
		prevList = mergedListHead
	}
	return prevList
}
