package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func (l *LinkedList) print() {
	head := l.head
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func (l *LinkedList) push(key int) {
	node := &Node{data: key}
	node.next = l.head
	l.head = node
}

func (l *LinkedList) delete(key int) {
	temp := l.head
	var prev *Node
	if temp != nil && temp.data == key {
		l.head = temp.next
		return
	}
	for temp != nil && temp.data != key {
		prev = temp
		temp = temp.next
	}
	if temp != nil {
		prev.next = temp.next
	}
}