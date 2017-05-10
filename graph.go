package kraken

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/satori/go.uuid"
)

// Graph holding the entire graph network.
type Graph struct {
	ID    uuid.UUID
	Name  string
	Nodes map[*Node]bool
}

// Inspect this graph.
func (g *Graph) Inspect() {
	fmt.Printf("ID:\t\t%s\n", g.ID)
	fmt.Printf("Type:\t\tGraph\n")
	fmt.Printf("Name:\t\t%s\n", g.Name)
	fmt.Printf("Size:\t\t%d\n", g.Size())
	fmt.Printf("Nodes:\t\t%d\n", g.CountNodes())
	fmt.Printf("\n")
}

// Size of this Graph struct.
func (g *Graph) Size() int {
	size := int(unsafe.Sizeof(g.ID))
	size = len(g.Name)
	for elem := range g.Nodes {
		size += elem.Size()
	}
	return size
}

// AddNode to a graph.
func (g *Graph) AddNode(n *Node) bool {
	_, found := g.Nodes[n]
	g.Nodes[n] = true
	return !found
}

// DeleteNode from a graph
func (g *Graph) DeleteNode(n *Node) {
	delete(g.Nodes, n)
}

// CountNodes returns the total number of nodes in the graph
func (g *Graph) CountNodes() int {
	return len(g.Nodes)
}

// GetNode tries to find a node based on an ID.
func (g *Graph) GetNode(id string) (n *Node, err error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	for elem := range g.Nodes {
		if elem.ID == uid {
			return elem, nil
		}
	}
	return nil, errors.New("no node found")
}

// SaveToDisk writes the content of this graph to disk.
func (g *Graph) SaveToDisk() bool {
	// TODO: Write to disk
	return false
}

// LoadFromDisk loads the graph from the disk.
// Needs the name of the graph to load.
func LoadFromDisk(name string) *Graph {
	// TODO: Load from Disk
	return nil
}

// NewGraph creates a brand new graph
func NewGraph(name string) *Graph {
	return &Graph{
		ID:    uuid.NewV4(),
		Name:  name,
		Nodes: make(map[*Node]bool),
	}
}
