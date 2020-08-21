package recursion

func MagicSlow(n []int) int {
	for i := range n {
		if n[i] == i {
			return i
		}
	}

	return -1
}
