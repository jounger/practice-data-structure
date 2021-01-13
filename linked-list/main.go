package main

import (
	"fmt"
)

func main() {
	// init a linked list
	list := &LinkedList{head: &Node{data: 1}}
	// push new data in list
	list.push(0)
	list.push(-1)
	// print out the list
	list.print()
	// delete node by given key
	list.delete(0)
	list.print()

	// delete node by given index
	list.deleteIdx(1)
	list.print()
	// check if single linked list is palindrome
	list.push(0)
	list.push(-1)
	list.print()
	fmt.Println(list.isPalindrome())
}