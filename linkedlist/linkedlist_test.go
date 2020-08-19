package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	cases := []struct {
		index int
		value string
	}{
		{0, "a"},
		{1, "b"},
		{2, "c"},
		{3, "d"},
	}

	list := New()
	for _, c := range cases {
		list.Add(&Node{Value: c.value})
	}

	for _, c := range cases {
		n, ok := list.Get(c.index)
		if !ok || n.Value != c.value {
			t.Errorf("get(%d): ok=%v, v=%v", c.index, ok, n)
		}
	}

	rcases := []struct {
		index int
		value string
	}{
		{2, "abd"},
		{1, "ad"},
		{0, "d"},
	}

	for _, c := range rcases {
		ok := list.Remove(c.index)
		if !ok {
			t.Errorf("remove(%d)", c.index)
		}

		if list.String() != c.value {
			t.Errorf("list=%s", list)
		}
	}

}
