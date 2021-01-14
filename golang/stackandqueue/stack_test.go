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
		if out := stack.pop(); out != val {
			t.Errorf("Stack.pop() != %d", val)
		}
	}
}