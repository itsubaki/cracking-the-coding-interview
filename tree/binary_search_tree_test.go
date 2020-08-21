package tree

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	// 3 -> 1, 4
	// 1 -> nil, 2
	// 4 -> nil, 5
	bst := NewMinBST([]int{1, 2, 3, 4, 5})
	if bst.Value != 3 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Left.Value != 1 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Right.Value != 4 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Left.Left != nil {
		t.Errorf("actual=%v", bst)
	}

	if bst.Left.Right.Value != 2 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Right.Left != nil {
		t.Errorf("actual=%v", bst)
	}

	if bst.Right.Right.Value != 5 {
		t.Errorf("actual=%v", bst)
	}
}
