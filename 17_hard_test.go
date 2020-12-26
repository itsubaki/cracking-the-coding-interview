package main_test

import "testing"

func TestAddWithoutPlus(t *testing.T) {
	Add := func(a, b int) int {
		for {
			if b == 0 {
				break
			}

			sum, carry := a^b, (a&b)<<1
			a, b = sum, carry
		}

		return a
	}

	cases := []struct {
		a, b, c int
	}{
		{10, 5, 15},
		{2, 15, 17},
	}

	for _, c := range cases {
		result := Add(c.a, c.b)
		if result == c.c {
			continue
		}

		t.Errorf("expected=%v, actual=%v", c.c, result)
	}
}
