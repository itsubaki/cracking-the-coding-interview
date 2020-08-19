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
	magicSlow := func(n []int) int {
		for i := range n {
			if n[i] == i {
				return i
			}
		}

		return -1
	}

	if magicSlow([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}) != 5 {
		t.Fail()
	}
}

func TestPowerSet(t *testing.T) {

}

func TestRecursiveMultiply(t *testing.T) {

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
