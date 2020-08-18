package queue

import "github.com/itsubaki/cracking-the-coding-interview/stack"

type MyQueue struct {
	Newest *stack.Stack
	Oldest *stack.Stack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		Newest: stack.New(),
		Oldest: stack.New(),
	}
}

func (q *MyQueue) shift() {
	for {
		if q.Newest.IsEmpty() {
			break
		}

		v := q.Newest.Pop()
		q.Oldest.Push(v)
	}
}

func (q *MyQueue) Add(value int) {
	q.Newest.Push(value)
}

func (q *MyQueue) Peek() int {
	q.shift()
	return q.Oldest.Peek()
}

func (q *MyQueue) Remove() int {
	q.shift()
	return q.Oldest.Pop()
}
