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

// Print out the list
func (l *LinkedList) print() {
	head := l.head
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

// Push in front of the list
func (l *LinkedList) push(key int) {
	node := &Node{data: key}
	node.next = l.head
	l.head = node
}

// Delete a node by given key
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

// Delete a node by given index
func (l *LinkedList) deleteIdx(idx int) {
	temp := l.head
	var prev *Node
	if temp != nil && idx == 0 {
		l.head = temp.next
		return
	}
	for i := 0; i < idx && temp != nil; i++ {
		prev = temp
		temp = temp.next
	}
	if temp != nil {
		prev.next = temp.next
	}
}