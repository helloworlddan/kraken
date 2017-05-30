package kraken

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
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
	Nodes    []*Node
}

// Inspect this graph.
func (g *Graph) Inspect() {
	fmt.Printf("ID:\t\t%s\n", g.ID)
	fmt.Printf("Type:\t\tGraph\n")
	fmt.Printf("Name:\t\t%s\n", g.Name)
	fmt.Printf("Created:\t%s\n", g.Created.Format(C.TimeFormat))
	fmt.Printf("Modified:\t%s\n", g.Modified.Format(C.TimeFormat))
	fmt.Printf("Saved:\t\t%s\n", g.Saved.Format(C.TimeFormat))
	fmt.Printf("Size:\t\t%d\n", g.Size())
	fmt.Printf("Nodes:\t\t%d\n", g.CountNodes())
	fmt.Printf("\n")
}

// Size of this Graph struct.
func (g *Graph) Size() (size int) {
	size = int(unsafe.Sizeof(g.ID))
	size = len(g.Name)
	for _, elem := range g.Nodes {
		size += elem.Size()
	}
	return size
}

// Update this Graph with another one.
func (g *Graph) Update(update *Graph) {
	g.Name = update.Name
	g.Nodes = update.Nodes
	g.Modified = time.Now()
}

// AddNode to a graph.
func (g *Graph) AddNode(n *Node) {
	index := -1
	for i, elem := range g.Nodes {
		if n == elem {
			index = i
		}
	}
	if index == -1 {
		g.Nodes = append(g.Nodes, n)
		g.Modified = time.Now()
	}
}

// DeleteNode from a graph
func (g *Graph) DeleteNode(n *Node) {
	index := -1
	for i, elem := range g.Nodes {
		if n == elem {
			index = i
		}
	}
	if index > -1 {
		g.Nodes = append(g.Nodes[:index], g.Nodes[index+1:]...)
		g.Modified = time.Now()
	}
}

// CountNodes returns the total number of nodes in the graph
func (g *Graph) CountNodes() (num int) {
	return len(g.Nodes)
}

// GetNode tries to find a node based on an ID.
func (g *Graph) GetNode(id string) (n *Node, err error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	for _, n = range g.Nodes {
		if n.ID == uid {
			return n, nil
		}
	}
	return nil, errors.New("node not found")
}

// GraphFromYaml recreates Graph from YAML
func GraphFromYaml(y string) (g *Graph, err error) {
	g = NewGraph("")
	err = yaml.Unmarshal([]byte(y), g)
	return g, err
}

// ToYaml transforms the content of this graph to yaml.
func (g *Graph) ToYaml() (out string, err error) {
	yam, err := yaml.Marshal(g)
	return string(yam), err
}

// GraphFromJSON recreates Graph from JSON
func GraphFromJSON(js string) (g *Graph, err error) {
	g = NewGraph("")
	err = json.Unmarshal([]byte(js), g)
	return g, err
}

// ToJSON transforms the content of this Engine to json..
func (g *Graph) ToJSON() (out string, err error) {
	js, err := json.Marshal(g)
	return string(js), err
}

// GraphFromXML recreates Graph from XML
func GraphFromXML(x string) (g *Graph, err error) {
	g = NewGraph("")
	err = xml.Unmarshal([]byte(x), g)
	return g, err
}

// ToXML transforms the content of this Engine to xml.
func (g *Graph) ToXML() (out string, err error) {
	x, err := xml.Marshal(g)
	return string(x), err
}

// DeserializeGraph this Graph.
func DeserializeGraph(raw string) (g *Graph, err error) {
	switch strings.ToUpper(C.OutputFormat) {
	case "YAML":
		return GraphFromYaml(raw)
	case "JSON":
		return GraphFromJSON(raw)
	case "XML":
		return GraphFromXML(raw)
	default:
		return nil, errors.New("Input format " + C.OutputFormat + " not recognized.")
	}
}

// Serialize this Graph.
func (g *Graph) Serialize() (out string, err error) {
	switch strings.ToUpper(C.OutputFormat) {
	case "YAML":
		return g.ToYaml()
	case "JSON":
		return g.ToJSON()
	case "XML":
		return g.ToXML()
	default:
		return "", errors.New("Output format " + C.OutputFormat + " not recognized.")
	}
}

// NewGraph creates a brand new graph
func NewGraph(name string) *Graph {
	return &Graph{
		Created:  time.Now(),
		ID:       uuid.NewV4(),
		Name:     name,
		Nodes:    make([]*Node, 0),
		Modified: time.Now(),
	}
}
