package linkedlist

import "strings"

type LinkedList struct {
	Size int
	Head *Node
}

func New() *LinkedList {
	return &LinkedList{
		Size: 0,
		Head: nil,
	}
}

func (l *LinkedList) Add(e *Node) {
	l.Size++
	if l.Head == nil {
		l.Head = e
		return
	}

	n := l.Head
	for {
		if n.Next == nil {
			n.Next = e
			return
		}

		n = n.Next
	}
}

func (l *LinkedList) Get(i int) (*Node, bool) {
	if i > l.Size-1 {
		return nil, false
	}

	if i == 0 {
		return l.Head, true
	}

	k, n := 1, l.Head
	for {
		if k == i {
			return n.Next, true
		}

		k++
		n = n.Next
	}
}

func (l *LinkedList) Remove(i int) bool {
	if i > l.Size-1 {
		return false
	}

	k, n := 0, l.Head
	for {
		if k == i {
			n.Value = n.Next.Value
			n.Next = n.Next.Next
			l.Size--
			return true
		}

		k++
		n = n.Next
	}
}

func (l *LinkedList) String() string {
	if l.Head == nil {
		return ""
	}

	var sb strings.Builder

	n := l.Head
	for {
		sb.WriteString(n.Value)
		if n.Next == nil {
			return sb.String()
		}

		n = n.Next
	}
}
