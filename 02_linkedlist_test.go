package main_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/linkedlist"
)

func TestRemoveDups(t *testing.T) {
	RemoveDups := func(list *linkedlist.LinkedList) *linkedlist.LinkedList {
		set := make(map[string]bool)

		n := list.Head
		prev := &linkedlist.Node{}
		for {
			if n == nil {
				break
			}

			if _, ok := set[n.Value]; ok {
				prev.Next = n.Next
			} else {
				set[n.Value] = true
				prev = n
			}

			n = n.Next
		}

		return list
	}

	list := linkedlist.New()
	for _, c := range []string{"a", "b", "a", "d", "a"} {
		list.Add(c)
	}

	result := RemoveDups(list)
	if result.String() != "abd" {
		t.Errorf("list=%s", result)
	}
}

func TestReturnKthToLast(t *testing.T) {
	Kth := func(list *linkedlist.LinkedList, k int) (string, bool) {
		p1, p2 := list.Head, list.Head

		for i := 0; i < k; i++ {
			if p1 == nil {
				return "", false
			}

			p1 = p1.Next
		}

		for {
			if p1.Next == nil {
				break
			}

			p1, p2 = p1.Next, p2.Next
		}

		return p2.Value, true
	}

	list := linkedlist.New()
	for _, c := range []string{"a", "b", "c", "d", "e"} {
		list.Add(c)
	}

	cases := []struct {
		k     int
		value string
	}{
		{0, "e"},
		{1, "d"},
		{2, "c"},
		{3, "b"},
		{4, "a"},
	}

	for _, c := range cases {
		result, ok := Kth(list, c.k)
		if ok && result == c.value {
			continue
		}

		t.Errorf("k=%v, ok=%v, expected=%v, actual=%v", c.k, ok, c.value, result)
	}
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
	Reverse := func(l *linkedlist.LinkedList) *linkedlist.LinkedList {
		ret := linkedlist.New()

		for i := l.Size - 1; i >= 0; i-- {
			v, ok := l.Get(i)
			if !ok {
				panic("ERROR")
			}

			ret.Add(v)
		}

		return ret
	}

	IsEqual := func(l1, l2 *linkedlist.LinkedList) bool {
		if l1.Size != l2.Size {
			return false
		}

		p1, p2 := l1.Head, l2.Head
		for {
			if p1 == nil && p2 == nil {
				break
			}

			if p1.Value != p2.Value {
				return false
			}

			p1, p2 = p1.Next, p2.Next
		}

		return true
	}

	Palindrome := func(list *linkedlist.LinkedList) bool {
		return IsEqual(list, Reverse(list))
	}

	list := linkedlist.New()
	for _, c := range []string{"a", "b", "c", "b", "a"} {
		list.Add(c)
	}

	if !Palindrome(list) {
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {

}

func TestLoopDetection(t *testing.T) {

}
