package recursion

func TripleStep(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return TripleStep(n-1) + TripleStep(n-2) + TripleStep(n-3)
}
