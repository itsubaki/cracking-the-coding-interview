package recursion

func MinProduct(a, b int) int {
	bigger, smaller := a, b
	if a < b {
		bigger, smaller = b, a
	}

	return minProductHelper(smaller, bigger)
}

func minProductHelper(smaller, bigger int) int {
	if smaller == 0 {
		return 0
	}

	if smaller == 1 {
		return bigger
	}

	s := smaller >> 1
	half := minProductHelper(s, bigger)

	if smaller%2 == 0 {
		return half + half
	}

	return half + half + bigger
}
