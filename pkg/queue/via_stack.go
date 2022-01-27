package queue

import "github.com/itsubaki/cracking-the-coding-interview/pkg/stack"

type Queue struct {
	in  *stack.Stack
	out *stack.Stack
}

func NewViaStack() *Queue {
	return &Queue{
		in:  stack.New(),
		out: stack.New(),
	}
}

func (q *Queue) Enq(e int) {
	q.in.Push(e)
}

func (q *Queue) Deq() int {
	if q.out.IsEmpty() {
		q.shift()
	}

	return q.out.Pop()
}

func (q *Queue) Front() int {
	if q.out.IsEmpty() {
		q.shift()
	}

	return q.out.Top()
}

func (q *Queue) shift() {
	for {
		if q.in.IsEmpty() {
			break
		}

		v := q.in.Pop()
		q.out.Push(v)
	}
}
