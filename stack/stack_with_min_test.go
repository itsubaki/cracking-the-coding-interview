package stack_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/stack"
)

func TestStackWithMin(t *testing.T) {
	s := stack.NewWithMin()

	push := []struct {
		value int
		min   int
	}{
		{5, 5},
		{6, 5},
		{3, 3},
		{7, 3},
	}

	for _, c := range push {
		s.Push(c.value)
		if s.Min() != c.min {
			t.Errorf("StackWithMin: push: min: expected=%v, actual=%v", c.min, s.Min())
		}
	}

	pop := []struct {
		value int
		min   int
	}{
		{7, 3},
		{3, 5},
		{6, 5},
		{5, 127},
	}

	for _, c := range pop {
		if p := s.Pop(); p != c.value {
			t.Errorf("StackWithMin: pop: expected=%v, actual=%v", c.value, p)
		}

		if s.Min() != c.min {
			t.Errorf("StackWithMin: pop: min: expected=%v, actual=%v", c.min, s.Min())
		}
	}
}
