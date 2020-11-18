package digraph

// New creates an empty, mutable directed graph.
func New() Graph {
	id := 0
	return &graph{
		id: &id,
	}
}

// Graph represents a directed graph.
type Graph interface {
	// NewNode creates a node inside the graph and returns it.
	NewNode() Node

	// Nodes returns all nodes.
	Nodes() []Node
}

type graph struct {
	id         *int
	lastNodeID int
	nodes      []Node
}

func (g *graph) NewNode() Node {
	g.lastNodeID++
	n := node{
		graphID: g.id,
		nodeID:  g.lastNodeID,
	}
	g.nodes = append(g.nodes, n)
	return n
}

func (g *graph) Nodes() []Node {
	return g.nodes
}

// Node is a node of a directed graph. It is a value object. Can be used as
// map key. It is also possible to compare nodes with ==.
type Node interface {
	internal()
}

type node struct {
	graphID *int
	nodeID  int
}

func (n node) internal() {}
