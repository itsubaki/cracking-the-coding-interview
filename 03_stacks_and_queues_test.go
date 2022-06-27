package main_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/queue"
	"github.com/itsubaki/cracking-the-coding-interview/stack"
)

func TestThreeInOne(t *testing.T) {

}

func TestStackMin(t *testing.T) {
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

func TestStackOfPlates(t *testing.T) {

}

func TestQueueViaStack(t *testing.T) {
	cases := []struct {
		value int
	}{
		{1},
		{2},
		{3},
		{4},
		{5},
	}

	q := queue.NewViaStack()
	for _, c := range cases {
		q.Enq(c.value)
	}

	for _, c := range cases {
		if q.Deq() == c.value {
			continue
		}
		t.Fail()
	}
}

func TestSortStack(t *testing.T) {

}

func TestAnimalShelter(t *testing.T) {

}
