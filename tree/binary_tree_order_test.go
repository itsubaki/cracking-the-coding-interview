package tree

import (
	"testing"
)

func TestInOrder(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	expected := []int{2, 4, 6, 8, 10, 20}
	actual := root.InOrder()
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

func TestPreOrder(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	expected := []int{8, 4, 2, 6, 10, 20}
	actual := root.PreOrder()
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

func TestPostOrder(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	expected := []int{2, 6, 4, 20, 10, 8}
	actual := root.PostOrder()
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
