package queue

import "testing"

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

	q := NewViaStack()
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
