package graph

type State int

const (
	Unvisited = "Unvisited"
	Visited   = "Visited"
	Visiting  = "Visiting"
)

type Node struct {
	Name     string
	State    string
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
