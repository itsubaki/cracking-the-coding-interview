package graph

type State int

const (
	Unvisited State = iota
	Visited
	Visiting
)

type Node struct {
	Name     string
	State    State
	Children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		State:    Unvisited,
		Children: make([]*Node, 0),
	}
}

func (n *Node) Add(c *Node) {
	n.Children = append(n.Children, c)
}
