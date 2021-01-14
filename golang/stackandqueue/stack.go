package main

import (
	"errors"
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
func (s *Stack) push(key int)  {
	if s.top > maxSize - 1 {
		errors.New("Stack overflow")
		return
	}
	s.top++
	s.data[s.top] = key
}

func (s *Stack) pop() int {
	if s.top < 0 {
		errors.New("Stack underflow")
		return 0
	}
	x := s.data[s.top]
	s.top--
	return x
}