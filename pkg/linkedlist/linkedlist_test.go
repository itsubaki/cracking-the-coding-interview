package linkedlist_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/pkg/linkedlist"
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

	list := linkedlist.New()
	for _, c := range cases {
		list.Add(c.value)
	}

	for _, c := range cases {
		v, ok := list.Get(c.index)
		if !ok || v != c.value {
			t.Errorf("get(%d): ok=%v, v=%v", c.index, ok, v)
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
