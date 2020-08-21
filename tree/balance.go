package tree

func IsBalanced(root *BinaryTree) bool {
	if root == nil {
		return true
	}

	diff := GetHeight(root.Left) - GetHeight(root.Right)
	if Abs(diff) > 1 {
		return false
	}

	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

func GetHeight(root *BinaryTree) int {
	if root == nil {
		return -1
	}

	return Max(GetHeight(root.Left), GetHeight(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
