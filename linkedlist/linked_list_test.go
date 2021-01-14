package main

import (
	"testing"
)

func TestPush(t *testing.T) {
	list := &LinkedList{}
	list.push(0)
	if list.head.data != 0 {
		t.Errorf("LinkedList.push(0) != %d", list.head.data)
	}
}

func TestAppend(t *testing.T) {
	list := &LinkedList{}
	list.push(0)
	list.append(1)
	for list.head != nil {
		if list.head.next == nil {
			if list.head.data != 1 {
				t.Errorf("LinkedList.append(0) != %d", list.head.data)
			}
			return
		}
		list.head = list.head.next
	}
}