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
		if head.next != nil {
			head = head.next
		} else {
			break
		}
	}
	fmt.Println()
}

// Print out the list reverse
func printReverse(node *Node) {
	if node == nil {
		return
	}
	printReverse(node.next) // func return void
	fmt.Print(node.data, " ")
}

// Push in front of the list
func (l *LinkedList) push(key int) {
	node := &Node{data: key}
	node.next = l.head
	l.head = node
}

// Append in end of the list
func (l *LinkedList) append(key int) {
	node := &Node{data: key}
	node.next = nil
	if l.head == nil {
		l.head = node
		return
	}
	for head := l.head; head != nil; head = head.next {
		if head.next == nil {
			head.next = node
			return
		}
	}
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

// Delete the list
func (l *LinkedList) destroy() {
	l.head = nil
}

// Count the length of list
func (l *LinkedList) length() int {
	head := l.head
	count := 0
	for head != nil {
		count++
		head = head.next
	}
	return count
}

// Find an element in list and return boolean
func (l *LinkedList) isExist(key int) bool {
	head := l.head
	for head != nil {
		if head.data == key {
			return true
		}
		head = head.next
	}
	return false
}

// Find an element in list by index and return it
func (l *LinkedList) get(idx int) *Node {
	head := l.head
	for i := 0; i < idx && head != nil; i++ {
		if i == idx {
			return head
		}
		head = head.next
	}
	return nil
}

// Count how many times a key show up in list
func (l *LinkedList) countTimes(key int) int {
	head := l.head
	count := 0
	for head != nil {
		if head.data == key {
			count++
		}
		head = head.next
	}
	return count
}

// Detect if a list has been loop
func (l *LinkedList) detectLoop() bool {
	m := make(map[*int]int)
	head := l.head
	for head != nil {
		if m[&head.data] > 1 {
			return true
		}
		m[&head.data] += 1 
		head = head.next
	}
	return false
}

// Check if single linked list is palindrome
func (l *LinkedList) isPalindrome() bool {
	stack := []int{}
	for head := l.head; head != nil; head = head.next {
		stack = append(stack, head.data)
	}
	i := len(stack) - 1
	for head := l.head; head != nil; head = head.next {
		if head.data == stack[i] {
			if i == 0 {
				return true
			}
			i--
			continue
		}
		return false
	}
	
	return false
}

// Remove duplicate element in list (sorted)
func (l *LinkedList) removeDuplicatesSorted() {
	head := l.head
	for head != nil {
		if head.next != nil && head.data == head.next.data {
			head.next = head.next.next
		} else {
			head = head.next
		}
	}
}

// Remove duplicate element in list (unsorted)
func (l *LinkedList) removeDuplicatesUnsorted() {
	m := make(map[int]int)
	temp := l.head
	var prev *Node

	for temp != nil {
		if m[temp.data] >= 1 {
			prev.next = temp.next
		} else {
			m[temp.data] += 1
			prev = temp
		}
		temp = temp.next
	}
}

// Remove duplicate element in list (unsorted)
func (l *LinkedList) reverse() {
	// TODO: reverse a list
	var prev, next *Node
	curr := l.head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}