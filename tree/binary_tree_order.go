package tree

import "fmt"

func InOrder(t *BinaryTree) {
	if t == nil {
		return
	}

	InOrder(t.Left)
	fmt.Println(t.Value)
	InOrder(t.Right)
}

func PreOrder(t *BinaryTree) {
	if t == nil {
		return
	}

	fmt.Println(t.Value)
	PreOrder(t.Left)
	PreOrder(t.Right)
}

func PostOrder(t *BinaryTree) {
	if t == nil {
		return
	}

	PostOrder(t.Left)
	PostOrder(t.Right)
	fmt.Println(t.Value)
}
