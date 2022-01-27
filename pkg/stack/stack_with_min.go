package stack

import "math"

type StackWithMin struct {
	values []int
	min    *Stack
}

func NewWithMin() *StackWithMin {
	return &StackWithMin{
		values: make([]int, 0),
		min: &Stack{
			values: make([]int, 0),
		},
	}

}

func (s *StackWithMin) Push(v int) {
	if v < s.Min() {
		s.min.Push(v)
	}

	s.values = append(s.values, v)
}

func (s *StackWithMin) Pop() int {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	if v == s.Min() {
		s.min.Pop()
	}

	return v
}

func (s *StackWithMin) Peek() int {
	return s.values[len(s.values)-1]
}

func (s *StackWithMin) Min() int {
	if s.min.IsEmpty() {
		return math.MaxInt8
	}

	return s.min.Top()
}
