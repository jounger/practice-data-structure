package main

import (
	"testing"
)

func TestEnQueue(t *testing.T) {
	var queue Queue
	queue.init()
	queue.enQueue(1)
	if queue.s1.data[0] != 1 {
		t.Errorf("Queue.enQueue(%d) != %d", 1, 1)
	}
}

func TestDeQueue(t *testing.T) {
	var queue Queue
	queue.init()
	queue.enQueue(1)
	queue.enQueue(2)
	queue.enQueue(3)
	out := queue.deQueue()
	if out != 1 {
		t.Errorf("Queue.deQueue() != %d", 1)
	}
}