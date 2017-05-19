package kraken

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

// Node in a graph.
type Node struct {
	iD       uuid.UUID
	created  time.Time
	modified time.Time
	name     string
	data     map[string]string
}

// ID gets the iD of this Node.
func (n *Node) ID() uuid.UUID {
	return n.iD
}

// Created gets the created of this Node.
func (n *Node) Created() time.Time {
	return n.created
}

// Modified gets the modified of this Node.
func (n *Node) Modified() time.Time {
	return n.modified
}

// Name gets the name of this Node.
func (n *Node) Name() string {
	return n.name
}

// Inspect this node.
func (n *Node) Inspect() {
	fmt.Printf("ID:\t\t%s\n", n.ID())
	fmt.Printf("Type:\t\tNode\n")
	fmt.Printf("Name:\t\t%s\n", n.Name())
	fmt.Printf("Created:\t%s\n", n.Created().Format(C.TimeFormat()))
	fmt.Printf("Modified:\t%s\n", n.Modified().Format(C.TimeFormat()))
	fmt.Printf("Size:\t\t%d\n", n.Size())
	fmt.Printf("Data:\n")
	for k, v := range n.data {
		fmt.Printf("\t%s => %s\n", k, v)
	}
	fmt.Printf("\n")
}

// Size of this Node struct.
func (n *Node) Size() int {
	size := int(unsafe.Sizeof(n.ID))
	size += len(n.Name())
	for k, v := range n.data {
		size += len(k)
		size += len(v)
	}
	return size
}

// PutData into a Node. Will always modify.
func (n *Node) PutData(key string, value string) {
	n.data[key] = value
	n.modified = time.Now()
}

// DropData from a Node. Do nothing if the item is not found.
func (n *Node) DropData(key string) {
	for k := range n.data {
		if k == key {
			delete(n.data, key)
			n.modified = time.Now()
		}
	}
}

// CountData returns the total number of data items in a Node.
func (n *Node) CountData() int {
	return len(n.data)
}

// FindData tries to find a data item based on its key.
func (n *Node) FindData(key string) (value string, err error) {
	for k, v := range n.data {
		if k == key {
			return v, nil
		}
	}
	return "", errors.New("key not found")
}

// ToYaml transforms the content of this Node to yaml.
func (n *Node) ToYaml() (y string, e error) {
	yam, err := yaml.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(yam), nil
}

// NewNode creates a brand new node
func NewNode(name string) *Node {
	return &Node{
		created:  time.Now(),
		iD:       uuid.NewV4(),
		name:     name,
		data:     make(map[string]string),
		modified: time.Now(),
	}
}
