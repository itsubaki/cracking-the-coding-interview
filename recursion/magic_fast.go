package recursion

func MagicFast(array []int) int {
	return magicfast(array, 0, len(array)-1)
}

func magicfast(array []int, start, end int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2
	if array[mid] == mid {
		return mid
	}

	if array[mid] > mid {
		return magicfast(array, start, mid-1)
	}

	return magicfast(array, mid+1, end)
}
