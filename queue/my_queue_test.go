package queue

import "testing"

func TestMyQueue(t *testing.T) {
	q := NewMyQueue()
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		q.Add(v)
	}

	if q.Peek() != 1 {
		t.Fail()
	}

	if q.Remove() != 1 {
		t.Fail()
	}

	if q.Remove() != 2 {
		t.Fail()
	}
}
