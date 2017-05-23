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

// Node in a graph.
type Node struct {
	ID       uuid.UUID
	Created  time.Time
	Modified time.Time
	Name     string
	Data     Data
}

type Data struct {
	Map map[string]string
}

// Inspect this node.
func (n *Node) Inspect() {
	fmt.Printf("ID:\t\t%s\n", n.ID)
	fmt.Printf("Type:\t\tNode\n")
	fmt.Printf("Name:\t\t%s\n", n.Name)
	fmt.Printf("Created:\t%s\n", n.Created.Format(C.TimeFormat))
	fmt.Printf("Modified:\t%s\n", n.Modified.Format(C.TimeFormat))
	fmt.Printf("Size:\t\t%d\n", n.Size())
	fmt.Printf("Data:\n")
	for k, v := range n.Data.Map {
		fmt.Printf("\t%s => %s\n", k, v)
	}
	fmt.Printf("\n")
}

// Size of this Node struct.
func (n *Node) Size() int {
	size := int(unsafe.Sizeof(n.ID))
	size += len(n.Name)
	for k, v := range n.Data.Map {
		size += len(k)
		size += len(v)
	}
	return size
}

// PutData into a Node. Will always modify.
func (n *Node) PutData(key string, value string) {
	n.Data.Map[key] = value
	n.Modified = time.Now()
}

// DropData from a Node. Do nothing if the item is not found.
func (n *Node) DropData(key string) {
	for k := range n.Data.Map {
		if k == key {
			delete(n.Data.Map, key)
			n.Modified = time.Now()
		}
	}
}

// CountData returns the total number of data items in a Node.
func (n *Node) CountData() int {
	return len(n.Data.Map)
}

// FindData tries to find a data item based on its key.
func (n *Node) FindData(key string) (value string, err error) {
	for k, v := range n.Data.Map {
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

// ToJSON transforms the content of this Node to yaml.
func (n *Node) ToJSON() (string, error) {
	js, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

// ToXML transforms the content of this Engine to xml.
func (n *Node) ToXML() (string, error) {
	x, err := xml.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(x), nil
}

// MarshalXML marshalls the data map to XML
func (d *Data) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}

	for key, value := range d.Map {
		t := xml.StartElement{Name: xml.Name{Space: "", Local: key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{Name: t.Name})
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

// Serialize this Node.
func (n *Node) Serialize() (string, error) {
	switch strings.ToUpper(C.OutputFormat) {
	case "YAML":
		return n.ToYaml()
	case "JSON":
		return n.ToJSON()
	case "XML":
		return n.ToXML()
	default:
		return "", errors.New("Output format " + C.OutputFormat + " not recognized.")
	}
}

// NewNode creates a brand new node
func NewNode(name string) *Node {
	return &Node{
		Created:  time.Now(),
		ID:       uuid.NewV4(),
		Name:     name,
		Data:     Data{Map: make(map[string]string)},
		Modified: time.Now(),
	}
}
