package kraken

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	fmt.Printf("Created:\t%s\n", g.Created.Format(TimeFormat))
	fmt.Printf("Modified:\t%s\n", g.Modified.Format(TimeFormat))
	fmt.Printf("Saved:\t\t%s\n", g.Saved.Format(TimeFormat))
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
	if !found {
		g.Modified = time.Now()
	}
	return !found
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
	return nil, errors.New("no node found")
}

// SaveToDisk writes the content of this graph to disk.
func (g *Graph) SaveToDisk() (err error) {
	g.Saved = time.Now()
	fileName := g.Name + ".kraken"

	y, err := yaml.Marshal(g)
	if err != nil {
		return err
	}
	data := []byte(string(y))

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadFromDisk loads the graph from the disk.
// Needs the name of the graph to load.
func LoadFromDisk(name string) (g *Graph, err error) {
	fileName := name + ".kraken"

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var gra Graph
	err = yaml.Unmarshal(data, &gra)
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
