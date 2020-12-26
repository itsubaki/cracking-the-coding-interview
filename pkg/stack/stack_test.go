package stack_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/pkg/stack"
)

func TestStack(t *testing.T) {
	cases := []struct {
		value int
	}{
		{5},
		{6},
		{3},
		{7},
	}

	s := stack.New()
	for _, c := range cases {
		s.Push(c.value)
	}

	for i := len(cases) - 1; i < 0; i-- {
		if s.Pop() == cases[i].value {
			continue
		}
		t.Errorf("%v", i)
	}
}
