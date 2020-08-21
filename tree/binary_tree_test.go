package tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	if root.Left.Value != 4 {
		t.Fail()
	}

	if root.Left.Left.Value != 2 {
		t.Fail()
	}

	if root.Left.Right.Value != 6 {
		t.Fail()
	}

	if root.Right.Value != 10 {
		t.Fail()
	}

	if root.Right.Right.Value != 20 {
		t.Fail()
	}

	if root.Right.Left != nil {
		t.Fail()
	}
}
