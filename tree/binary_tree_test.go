package tree_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/tree"
)

func TestBinaryTree(t *testing.T) {
	n4 := &tree.BinaryTree{Value: 4}
	n5 := &tree.BinaryTree{Value: 5}
	n2 := &tree.BinaryTree{Value: 2, Left: n4, Right: n5}
	n3 := &tree.BinaryTree{Value: 3}
	n1 := &tree.BinaryTree{Value: 1, Left: n2, Right: n3}

	{
		expected := []int{4, 2, 5, 1, 3}
		actual := n1.InOrder()
		if len(expected) != len(actual) {
			t.Errorf("expected=%v, actual=%v", expected, actual)
		}

		for i := range expected {
			if expected[i] == actual[i] {
				continue
			}

			t.Errorf("expected=%v, actual=%v", expected, actual)
		}
	}

	{
		expected := []int{1, 2, 4, 5, 3}
		actual := n1.PreOrder()
		if len(expected) != len(actual) {
			t.Errorf("expected=%v, actual=%v", expected, actual)
		}

		for i := range expected {
			if expected[i] == actual[i] {
				continue
			}

			t.Errorf("expected=%v, actual=%v", expected, actual)
		}
	}

	{
		expected := []int{4, 5, 2, 3, 1}
		actual := n1.PostOrder()
		if len(expected) != len(actual) {
			t.Errorf("expected=%v, actual=%v", expected, actual)
		}

		for i := range expected {
			if expected[i] == actual[i] {
				continue
			}

			t.Errorf("expected=%v, actual=%v", expected, actual)
		}
	}
}
