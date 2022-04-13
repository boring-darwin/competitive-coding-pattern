package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {

	node6 := &node{data: 6, next: nil}
	node5 := &node{data: 4, next: node6}
	node4 := &node{data: 4, next: node5}
	node3 := &node{data: 3, next: node4}
	node2 := &node{data: 2, next: node3}
	node1 := &node{data: 1, next: node2}

	start := node1
	print(start)
	findMiddle(start)

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

func rearrange(start *node) {

}

func findMiddle(start *node) *node {

	p1 := start
	p2 := start

	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next
	}

	fmt.Println(p1.data)

	return p1

}
