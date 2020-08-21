package tree

func NewMinBST(array []int) *BinaryTree {
	return newMinBST(array, 0, len(array)-1)
}

func newMinBST(array []int, start, end int) *BinaryTree {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	n := &BinaryTree{Value: array[mid]}
	n.Left = newMinBST(array, start, mid-1)
	n.Right = newMinBST(array, mid+1, end)

	return n
}
