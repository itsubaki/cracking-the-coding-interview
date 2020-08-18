package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	for _, i := range []int{5, 6, 3, 7} {
		s.Push(i)
	}

	for _, i := range []int{7, 3, 6, 5} {
		if s.Pop() != i {
			t.Errorf("%v", i)
		}
	}
}
