package main

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/linkedlist"
)

func TestRemoveDups(t *testing.T) {

}

func TestReturnKthToLast(t *testing.T) {

}

func TestDeleteMiddleNode(t *testing.T) {
	list := linkedlist.New()
	for _, c := range []string{"a", "b", "c", "d", "e"} {
		list.Add(c)
	}

	cases := []struct {
		index int
		ok    bool
		value string
	}{
		{1, true, "acde"},
		{2, true, "ace"},
		{1, true, "ae"},
	}

	for _, c := range cases {
		ok := list.Remove(c.index)
		if ok != c.ok {
			t.Errorf("remove(%d)", c.index)
		}

		if list.String() != c.value {
			t.Errorf("list=%s", list)
		}
	}

}

func TestPartition(t *testing.T) {

}

func TestSumLists(t *testing.T) {

}

func TestLinkedListPalindrome(t *testing.T) {

}

func TestIntersection(t *testing.T) {

}

func TestLoopDetection(t *testing.T) {

}
