package kraken

import (
	"errors"
	"fmt"
	"time"
	"unsafe"

	yaml "gopkg.in/yaml.v2"

	"github.com/satori/go.uuid"
)

// Node in a graph.
type Node struct {
	ID       uuid.UUID
	Created  time.Time
	Modified time.Time
	Name     string
	Data     map[string]string
}

// Inspect this node.
func (n *Node) Inspect() {
	fmt.Printf("ID:\t\t%s\n", n.ID)
	fmt.Printf("Type:\t\tNode\n")
	fmt.Printf("Name:\t\t%s\n", n.Name)
	fmt.Printf("Created:\t%s\n", n.Created.Format(TimeFormat))
	fmt.Printf("Modified:\t%s\n", n.Modified.Format(TimeFormat))
	fmt.Printf("Size:\t\t%d\n", n.Size())
	fmt.Printf("Data:\n")
	for k, v := range n.Data {
		fmt.Printf("\t%s => %s\n", k, v)
	}
	fmt.Printf("\n")
}

// Size of this Node struct.
func (n *Node) Size() int {
	size := int(unsafe.Sizeof(n.ID))
	size += len(n.Name)
	for k, v := range n.Data {
		size += len(k)
		size += len(v)
	}
	return size
}

// PutData into a Node. Will always modify.
func (n *Node) PutData(key string, value string) {
	n.Data[key] = value
	n.Modified = time.Now()
}

// DropData from a Node. Do nothing if the item is not found.
func (n *Node) DropData(key string) {
	for k := range n.Data {
		if k == key {
			delete(n.Data, key)
			n.Modified = time.Now()
		}
	}
}

// CountData returns the total number of data items in a Node.
func (n *Node) CountData() int {
	return len(n.Data)
}

// FindData tries to find a data item based on its key.
func (n *Node) FindData(key string) (value string, err error) {
	for k, v := range n.Data {
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
		Created:  time.Now(),
		ID:       uuid.NewV4(),
		Name:     name,
		Data:     make(map[string]string),
		Modified: time.Now(),
	}
}
