package main

import (
	"testing"
)

const maxInput int = 100

type Test struct {
	input [maxInput]int
	expected [maxInput]int
}

func TestPush(t *testing.T) {
	var list LinkedList
	test := &Test{
		input: [maxInput]int{1,2,3},
		expected: [maxInput]int{1,2,3},
	}
	for i, val := range test.input {
		list.push(val)
		if list.head.data != test.expected[i] {
			t.Errorf("LinkedList.push(%d) != %d", val, test.expected[i])
		}
	}	
}

func TestAppend(t *testing.T) {
	var list LinkedList
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

func TestDelete(t *testing.T) {

}