package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {

	var n5 *ListNode
	n4 := &ListNode{Val: 98, Next: n5}
	n3 := &ListNode{Val: 79, Next: n4}
	n2 := &ListNode{Val: 68, Next: n3}
	n1 := &ListNode{Val: 39, Next: n2}
	n5 = &ListNode{Val: 98, Next: n2}
	head1 := &ListNode{Val: 10, Next: n1}
	n4.Next = n5
	output := findCycle(head1)
	fmt.Println(output)
}

func findCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
