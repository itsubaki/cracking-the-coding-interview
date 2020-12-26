package tree

func InOrder(t *BinaryTree, out []int) []int {
	if t == nil {
		return out
	}

	l := InOrder(t.Left, out)
	l = append(l, t.Value)
	return InOrder(t.Right, l)
}

func PreOrder(t *BinaryTree, out []int) []int {
	if t == nil {
		return out
	}

	out = append(out, t.Value)
	l := PreOrder(t.Left, out)
	return PreOrder(t.Right, l)
}

func PostOrder(t *BinaryTree, out []int) []int {
	if t == nil {
		return out
	}

	l := PostOrder(t.Left, out)
	r := PostOrder(t.Right, l)
	return append(r, t.Value)
}
