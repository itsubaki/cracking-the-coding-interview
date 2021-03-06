package main_test

import (
	"math"
	"strings"
	"testing"
)

func TestInsertion(t *testing.T) {
	Update := func(n, m, i, j uint) uint {
		ones := math.MaxUint32
		left := uint(ones << (j + 1))
		right := uint(((1 << i) - 1))

		mask := left | right
		nCleared := n & mask
		mShifted := m << i

		return nCleared | mShifted
	}

	cases := []struct {
		n, m, i, j, result uint
	}{
		{1024, 19, 2, 6, 1100},
	}

	for _, c := range cases {
		result := Update(c.n, c.m, c.i, c.j)
		if result == c.result {
			continue
		}

		t.Errorf("Update: expected=%v, actual=%v", c.result, result)
	}
}

func TestBinaryToString(t *testing.T) {
	ToString := func(num float32) string {
		if num > 1 || num < 0 {
			return "ERROR"
		}

		var sb strings.Builder
		sb.WriteString(".")
		for {
			if num <= 0 {
				break
			}

			if sb.Len() > 31 {
				return "ERROR"
			}

			r := num * 2
			if r < 1 {
				sb.WriteString("0")
				num = r
				continue
			}

			sb.WriteString("1")
			num = r - 1
		}

		return sb.String()
	}

	cases := []struct {
		num float32
		bin string
	}{
		{0.125, ".001"},
		{0.5, ".1"},
		{0.7, ".101100110011001100110011"},
	}

	for _, c := range cases {
		result := ToString(c.num)
		if result == c.bin {
			continue
		}

		t.Errorf("ToString: expected=%v, actual=%v", c.bin, result)
	}
}

func TestConversion(t *testing.T) {
	BitSwapRequired := func(a, b int) int {
		count := 0
		for c := a ^ b; c != 0; c = c >> 1 {
			count = count + c&1
		}

		return count
	}

	cases := []struct {
		a, b, c int
	}{
		{29, 15, 2},
	}

	for _, c := range cases {
		result := BitSwapRequired(c.a, c.b)
		if result == c.c {
			continue
		}

		t.Errorf("BitSwapRequired: expected=%v, actual=%v", c.c, result)
	}
}

func TestPairwiseSwap(t *testing.T) {
	SwapOddEvenBits := func(x int) int {
		return ((x & 0xaaaaaaaa) >> 1) | ((x & 0x55555555) << 1)
	}

	cases := []struct {
		a, b int
	}{
		{10, 5}, // 10 (1010) -> 5 (0101)
	}

	for _, c := range cases {
		result := SwapOddEvenBits(c.a)
		if result != c.b {
			t.Errorf("expected=%v, actual=%v", c.b, result)
		}
	}
}
