package tree

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (t *BinaryTree) IsComplete() bool {
	// TODO
	return false
}

func (t *BinaryTree) IsFull() bool {
	// TODO
	return false
}

func (t *BinaryTree) IsPerfect() bool {
	return t.IsComplete() && t.IsFull()
}

func (t *BinaryTree) InOrder() []int {
	return InOrder(t, []int{})
}

func (t *BinaryTree) PreOrder() []int {
	return PreOrder(t, []int{})
}

func (t *BinaryTree) PostOrder() []int {
	return PostOrder(t, []int{})
}
