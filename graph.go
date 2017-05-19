package kraken

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

// Graph holding the entire graph network.
type Graph struct {
	ID       uuid.UUID
	Name     string
	Created  time.Time
	Modified time.Time
	Saved    time.Time
	Nodes    map[*Node]bool
}

// Inspect this graph.
func (g *Graph) Inspect() {
	fmt.Printf("ID:\t\t%s\n", g.ID)
	fmt.Printf("Type:\t\tGraph\n")
	fmt.Printf("Name:\t\t%s\n", g.Name)
	fmt.Printf("Created:\t%s\n", g.Created.Format(C.TimeFormat()))
	fmt.Printf("Modified:\t%s\n", g.Modified.Format(C.TimeFormat()))
	fmt.Printf("Saved:\t\t%s\n", g.Saved.Format(C.TimeFormat()))
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
func (g *Graph) AddNode(n *Node) {
	_, found := g.Nodes[n]
	g.Nodes[n] = true
	if !found {
		g.Modified = time.Now()
	}
}

// DeleteNode from a graph
func (g *Graph) DeleteNode(n *Node) {
	_, found := g.Nodes[n]
	delete(g.Nodes, n)
	if found {
		g.Modified = time.Now()
	}
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
	return nil, errors.New("node not found")
}

// ToYaml transforms the content of this graph to yaml.
func (g *Graph) ToYaml() (y string, e error) {
	yam, err := yaml.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(yam), nil
}

// FromYaml recreates Graph from YAML
func FromYaml(y string) (g *Graph, e error) {
	var gra Graph
	err := yaml.Unmarshal([]byte(y), &gra)
	if err != nil {
		return nil, err
	}
	return &gra, nil
}

// NewGraph creates a brand new graph
func NewGraph(name string) *Graph {
	return &Graph{
		Created:  time.Now(),
		ID:       uuid.NewV4(),
		Name:     name,
		Nodes:    make(map[*Node]bool),
		Modified: time.Now(),
	}
}
