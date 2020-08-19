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
		list.Add(c.value)
	}

	for _, c := range cases {
		n, ok := list.Get(c.index)
		if !ok || n.Value != c.value {
			t.Errorf("get(%d): ok=%v, v=%v", c.index, ok, n)
		}
	}

	rcases := []struct {
		index int
		ok    bool
		value string
	}{
		{2, true, "abd"},
		{1, true, "ad"},
		{0, true, "d"},
		{0, false, "d"},
	}

	for _, c := range rcases {
		ok := list.Remove(c.index)
		if ok != c.ok {
			t.Errorf("remove(%d)", c.index)
		}

		if list.String() != c.value {
			t.Errorf("list=%s", list)
		}
	}

}
