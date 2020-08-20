package main

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/graph"
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
