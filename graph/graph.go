package graph

type Graph struct {
	Nodes []*Node
}

func New() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
	}
}

func (g *Graph) Add(n *Node) {
	g.Nodes = append(g.Nodes, n)
}
