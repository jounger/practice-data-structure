package main

import (
	"errors"
	"fmt"
)

const maxSize int = 100

type Stack struct {
	top int
	data [maxSize]int
}

func (s *Stack) init(){
	s.top = -1
}

// push a key into stack
func (s *Stack) push(key int) error {
	if s.top > maxSize - 1 {
		return errors.New("Stack overflow")
	}
	s.top++
	s.data[s.top] = key
	return nil
}

func (s *Stack) pop() (int, error) {
	if s.top < 0 {
		err := errors.New("Stack underflow")
		return 0, err
	}
	x := s.data[s.top]
	s.top--
	return x, nil
}

func (s *Stack) isEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *Stack) isFull() bool {
	if s.top >= maxSize - 1 {
		return true
	}
	return false
}

func (s *Stack) peek(idx int) (int, error) {
	if idx > s.top {
		err := errors.New("Not found")
		return 0, err
	}
	return s.data[idx], nil
} 

func (s *Stack) count() int {
	if s.top < 0 {
		return 0
	}
	return s.top + 1
}

func (s *Stack) change(idx, val int) {
	s.data[idx] = val
}

func (s *Stack) display() {
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Print(s.data[i], " ")
	}
	fmt.Println()
}