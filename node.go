package kraken

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/satori/go.uuid"
)

// Node in a graph.
type Node struct {
	ID       uuid.UUID
	Created  time.Time
	Modified time.Time
	Name     string
}

// Inspect this node.
func (n *Node) Inspect() {
	fmt.Printf("ID:\t\t%s\n", n.ID)
	fmt.Printf("Type:\t\tNode\n")
	fmt.Printf("Name:\t\t%s\n", n.Name)
	fmt.Printf("Created:\t%s\n", n.Created.Format(time.RFC3339))
	fmt.Printf("Modified:\t%s\n", n.Modified.Format(time.RFC3339))
	fmt.Printf("Size:\t\t%d\n", n.Size())
	fmt.Printf("\n")
}

// Size of this Node struct.
func (n *Node) Size() int {
	size := int(unsafe.Sizeof(n.ID))
	size += len(n.Name)
	return size
}

// NewNode creates a brand new node
func NewNode(name string) *Node {
	return &Node{
		Created:  time.Now(),
		ID:       uuid.NewV4(),
		Name:     name,
		Modified: time.Now(),
	}
}
