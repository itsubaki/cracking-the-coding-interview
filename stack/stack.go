package stack

type Stack struct {
	values []int
}

func New() *Stack {
	return &Stack{
		values: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	ret := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return ret
}

func (s *Stack) Peek() int {
	return s.values[len(s.values)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}
