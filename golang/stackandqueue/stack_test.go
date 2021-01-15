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
	var stack Stack
	stack.init()
	test := &Test{
		input: [maxInput]int{1,2,3},
		expected: [maxInput]int{1,2,3},
	}

	for i, val := range test.input {
		stack.push(val)
		if stack.data[i] != val {
			t.Errorf("Stack.push(%d) != %d", val, val)
		}
	}
}

func TestPop(t *testing.T) {
	var stack Stack
	stack.init()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	for _, val := range []int{3,2,1} {
		if out, _ := stack.pop(); out != val {
			t.Errorf("Stack.pop() != %d", val)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var stack Stack
	stack.init()
	stack.push(1)

	if stack.isEmpty() {
		t.Errorf("Stack.isEmpty() != %v", false)
	}
	stack.pop()
	if !stack.isEmpty() {
		t.Errorf("Stack.isEmpty() != %v", true)
	}
}

func TestIsFull(t *testing.T) {
	var stack Stack
	stack.init()
	if stack.isFull() {
		t.Errorf("Stack.isEmpty() != %v", false)
	}
	for i := 0; i < 100; i++ {
		stack.push(i)
	}
	if !stack.isFull() {
		t.Errorf("Stack.isEmpty() != %v", true)
	}
}

func TestPeek(t *testing.T) {
	var stack Stack
	stack.init()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	p, _ := stack.peek(1)
	if p != 2 {
		t.Errorf("Stack.peek(%d) != %d", 1, 2)
	}
}

func TestCount(t *testing.T) {
	var stack Stack
	stack.init()
	stack.push(1)
	count := stack.count()
	if count != 1 {
		t.Errorf("Stack.count() != %d", 1)
	}
}

func TestChange(t *testing.T) {
	var stack Stack
	stack.init()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.change(2, 5)
	if out, _ := stack.peek(2); out != 5 {
		t.Errorf("Stack.change(%d, %d) != %d", 2, 5, 5)
	}
}