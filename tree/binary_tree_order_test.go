package tree

import "testing"

func TestInOrder(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	InOrder(root)
}

func TestPreOrder(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	PreOrder(root)
}

func TestPostOrder(t *testing.T) {
	root := &BinaryTree{Value: 8}
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(20)

	PostOrder(root)
}
