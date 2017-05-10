package kraken

import (
	"fmt"
	"math/rand"
)

// Node in a graph.
type Node struct {
	ID   int64
	Name string
}

// Inspect this node.
func (n *Node) Inspect() {
	fmt.Printf("ID:\t\t%d\n", n.ID)
	fmt.Printf("Type:\t\tNode\n")
	fmt.Printf("Name:\t\t%s\n", n.Name)
	fmt.Printf("Size:\t\t%d\n", n.Size())
	fmt.Printf("\n")
}

// Size of this Node struct.
func (n *Node) Size() int {
	// TODO: Compute size
	size := -1
	return size
}

// NewNode creates a brand new node
func NewNode(name string) *Node {
	return &Node{
		ID:   rand.Int63(),
		Name: name,
	}
}
