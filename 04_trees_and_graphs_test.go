package main

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/graph"
	"github.com/itsubaki/cracking-the-coding-interview/tree"
)

func TestRouteBetweenNodes(t *testing.T) {
	Search := func(g *graph.Graph, start, end *graph.Node) bool {
		if start == end {
			return true
		}

		for i := range g.Nodes {
			g.Nodes[i].State = graph.Unvisited
		}
		start.State = graph.Visiting

		// queue
		q := make([]*graph.Node, 0)
		q = append(q, start)

		u := q[0]
		for {
			if len(q) == 0 {
				break
			}

			u, q = q[0], q[1:]
			for i := range u.Children {
				v := u.Children[i]
				if v.State == graph.Unvisited {
					if v == end {
						return true
					}

					v.State = graph.Visiting
					q = append(q, v)
				}
			}

			u.State = graph.Visited
		}

		return false
	}

	start := graph.NewNode("start")
	m1 := graph.NewNode("m1")
	m2 := graph.NewNode("m2")
	end := graph.NewNode("end")

	end.Add(start)
	m2.Add(end)
	m1.Add(m2)
	start.Add(m1)

	g := graph.New()
	g.Add(start)

	if !Search(g, start, end) {
		t.Fail()
	}
}

func TestMinimalTree(t *testing.T) {
	// 3 -> 1, 4
	// 1 -> nil, 2
	// 4 -> nil, 5
	bst := tree.NewMinBST([]int{1, 2, 3, 4, 5})
	if bst.Value != 3 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Left.Value != 1 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Right.Value != 4 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Left.Left != nil {
		t.Errorf("actual=%v", bst)
	}

	if bst.Left.Right.Value != 2 {
		t.Errorf("actual=%v", bst)
	}

	if bst.Right.Left != nil {
		t.Errorf("actual=%v", bst)
	}

	if bst.Right.Right.Value != 5 {
		t.Errorf("actual=%v", bst)
	}
}

func TestValidateBST(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 7, 8, 9}
	actual := tree.NewMinBST(expected).InOrder()

	if len(expected) != len(actual) {
		t.Errorf("expected=%v, actual=%v", expected, actual)
	}

	for i := range expected {
		if expected[i] == actual[i] {
			continue
		}

		t.Errorf("expected=%v, actual=%v", expected, actual)
	}
}
