package main

import "testing"

func TestNumberSwapper(t *testing.T) {
	Swap := func(a, b int) (int, int) {
		a = a ^ b
		b = a ^ b
		a = a ^ b
		return a, b
	}

	cases := []struct {
		a, b int
	}{
		{10, 5},
		{2, 15},
	}

	for _, c := range cases {
		s1, s2 := Swap(c.a, c.b)
		if s1 == c.b && s2 == c.a {
			continue
		}

		t.Errorf("expected=%v %v, actual=%v %v", c.a, c.b, s1, s2)
	}
}
