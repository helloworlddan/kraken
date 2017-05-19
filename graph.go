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
	iD       uuid.UUID
	name     string
	created  time.Time
	modified time.Time
	saved    time.Time
	nodes    map[*Node]bool
}

// ID gets the id of this Graph
func (g *Graph) ID() uuid.UUID {
	return g.iD
}

// Name gets the name of this Graph
func (g *Graph) Name() string {
	return g.name
}

// Created gets the created of this Graph
func (g *Graph) Created() time.Time {
	return g.created
}

// Modified gets the modified of this Graph
func (g *Graph) Modified() time.Time {
	return g.modified
}

// Saved gets the saved of this Graph
func (g *Graph) Saved() time.Time {
	return g.saved
}

// JustModified sets the modification time to now.
func (g *Graph) JustModified() {
	g.modified = time.Now()
}

// JustSaved sets the saving time to now.
func (g *Graph) JustSaved() {
	g.saved = time.Now()
}

// Inspect this graph.
func (g *Graph) Inspect() {
	fmt.Printf("ID:\t\t%s\n", g.ID())
	fmt.Printf("Type:\t\tGraph\n")
	fmt.Printf("Name:\t\t%s\n", g.Name())
	fmt.Printf("Created:\t%s\n", g.Created().Format(C.TimeFormat()))
	fmt.Printf("Modified:\t%s\n", g.Modified().Format(C.TimeFormat()))
	fmt.Printf("Saved:\t\t%s\n", g.Saved().Format(C.TimeFormat()))
	fmt.Printf("Size:\t\t%d\n", g.Size())
	fmt.Printf("Nodes:\t\t%d\n", g.CountNodes())
	fmt.Printf("\n")
}

// Size of this Graph struct.
func (g *Graph) Size() int {
	size := int(unsafe.Sizeof(g.ID))
	size = len(g.Name())
	for elem := range g.nodes {
		size += elem.Size()
	}
	return size
}

// AddNode to a graph.
func (g *Graph) AddNode(n *Node) {
	_, found := g.nodes[n]
	g.nodes[n] = true
	if !found {
		g.modified = time.Now()
	}
}

// DeleteNode from a graph
func (g *Graph) DeleteNode(n *Node) {
	_, found := g.nodes[n]
	delete(g.nodes, n)
	if found {
		g.modified = time.Now()
	}
}

// CountNodes returns the total number of nodes in the graph
func (g *Graph) CountNodes() int {
	return len(g.nodes)
}

// GetNode tries to find a node based on an ID.
func (g *Graph) GetNode(id string) (n *Node, err error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	for elem := range g.nodes {
		if elem.ID() == uid {
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
		created:  time.Now(),
		iD:       uuid.NewV4(),
		name:     name,
		nodes:    make(map[*Node]bool),
		modified: time.Now(),
	}
}
