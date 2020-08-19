package queue

import "github.com/itsubaki/cracking-the-coding-interview/stack"

type ViaStack struct {
	Newest *stack.Stack
	Oldest *stack.Stack
}

func NewViaStack() *ViaStack {
	return &ViaStack{
		Newest: stack.New(),
		Oldest: stack.New(),
	}
}

func (q *ViaStack) shift() {
	for {
		if q.Newest.IsEmpty() {
			break
		}

		v := q.Newest.Pop()
		q.Oldest.Push(v)
	}
}

func (q *ViaStack) Add(value int) {
	q.Newest.Push(value)
}

func (q *ViaStack) Peek() int {
	q.shift()
	return q.Oldest.Peek()
}

func (q *ViaStack) Remove() int {
	q.shift()
	return q.Oldest.Pop()
}
