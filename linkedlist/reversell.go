package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func main() {

	node4 := &node{data: 6, next: nil}
	node3 := &node{data: 5, next: node4}
	node2 := &node{data: 3, next: node3}
	node1 := &node{data: 1, next: node2}

	print(node1)
	println()
	print(reverse(node1))

}

func print(curr *node) {

	for curr != nil {
		fmt.Printf("%d", curr.data)
		curr = curr.next
		if curr != nil {
			fmt.Printf("->")
		}
	}
}

func reverse(curr *node) *node {

	prev := curr
	curr = curr.next
	// fmt.Println(prev.data)

	prev.next = nil

	for curr != nil {
		now := curr
		// fmt.Println(now.data)
		curr = curr.next
		now.next = prev
		prev = now
	}

	return prev
}
