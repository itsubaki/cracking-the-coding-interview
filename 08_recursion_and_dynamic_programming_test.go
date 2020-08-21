package main

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/recursion"
)

func TestTripleStep(t *testing.T) {
	cases := []struct {
		num    int
		result int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
	}

	for _, c := range cases {
		result := recursion.TripleStep(c.num)
		if result == c.result {
			continue
		}

		t.Errorf("TripleStep(\"%v\"): expected=%v, actual=%v", c.num, c.result, result)
	}
}

func TestRobotInAGrid(t *testing.T) {

}

func TestMagicIndex(t *testing.T) {
	cases := []struct {
		a []int
		b int
	}{
		{[]int{-40, -20, -1, 1, 2, 3, 5, 7, 9, 12, 13}, 7},
	}

	for _, c := range cases {
		result := recursion.MagicSlow(c.a)
		if result == c.b {
			continue
		}

		t.Errorf("expected=%v, actual=%v", c.b, result)
	}

	for _, c := range cases {
		result := recursion.MagicFast(c.a)
		if result == c.b {
			continue
		}

		t.Errorf("expected=%v, actual=%v", c.b, result)
	}
}

func TestPowerSet(t *testing.T) {

}

func TestRecursiveMultiply(t *testing.T) {
	cases := []struct {
		a, b, c int
	}{
		{8, 7, 56},
		{9, 2, 18},
	}

	for _, c := range cases {
		result := recursion.MinProduct(c.a, c.b)
		if result == c.c {
			continue
		}

		t.Errorf("expected=%v, actual=%v", c.b, result)
	}
}

func TestTowersOfHanoi(t *testing.T) {

}

func TestPermutationsWithoutDups(t *testing.T) {

}

func TestPermutationsWithDuplicates(t *testing.T) {

}

func TestParens(t *testing.T) {

}

func TestPaintFill(t *testing.T) {

}

func TestCoins(t *testing.T) {

}

func TestEightQueens(t *testing.T) {

}

func TestStackOfBoxes(t *testing.T) {

}

func TestBooleanEvaluation(t *testing.T) {

}
