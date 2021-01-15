package main

type Queue struct {
	s1 *Stack
	s2 *Stack
}

func (q *Queue) init() {
	var s1 Stack
	var s2 Stack
	s1.init()
	s2.init()
	q.s1 = &s1
	q.s2 = &s2
}

func (q *Queue) enQueue(key int) {
	q.s1.push(key)
}

func (q *Queue) deQueue() int {
	if q.s1.top == 0 {
		out, _ := q.s1.pop()
		return out
	}

	if q.s2.isEmpty() {
		for !q.s1.isEmpty() {
			x, _ := q.s1.pop()
			q.s2.push(x)
		}
	}
	out, _ := q.s2.pop()
	return out
}