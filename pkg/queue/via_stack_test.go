package queue_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/pkg/queue"
)

func TestViaStack(t *testing.T) {
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
		q.Add(c.value)
	}

	for _, c := range cases {
		if q.Remove() == c.value {
			continue
		}
		t.Fail()
	}
}
