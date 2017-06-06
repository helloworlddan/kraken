package core

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
	Data     map[string]string
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
	for k, v := range n.Data {
		fmt.Printf("\t%s => %s\n", k, v)
	}
	fmt.Printf("\n")
}

// Size of this Node struct.
func (n *Node) Size() (size int) {
	size = int(unsafe.Sizeof(n.ID))
	size += len(n.Name)
	for k, v := range n.Data {
		size += len(k)
		size += len(v)
	}
	return size
}

// Update this Node with another one.
func (n *Node) Update(update *Node) {
	n.Name = update.Name
	n.Data = update.Data
	n.Modified = time.Now()
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
func (n *Node) CountData() (num int) {
	return len(n.Data)
}

// FindData tries to find a data item based on its key.
func (n *Node) FindData(key string) (value string, err error) {
	for k, value := range n.Data {
		if k == key {
			return value, err
		}
	}
	return "", errors.New("key not found")
}

// NodeFromYaml recreates Node from YAML
func NodeFromYaml(y string) (n *Node, err error) {
	n = NewNode("")
	err = yaml.Unmarshal([]byte(y), n)
	return n, err
}

// ToYaml transforms the content of this Node to yaml.
func (n *Node) ToYaml() (out string, err error) {
	yam, err := yaml.Marshal(n)
	return string(yam), err
}

// NodeFromJSON recreates Node from JSON
func NodeFromJSON(js string) (n *Node, err error) {
	n = NewNode("")
	err = json.Unmarshal([]byte(js), n)
	return n, err
}

// ToJSON transforms the content of this Node to yaml.
func (n *Node) ToJSON() (out string, err error) {
	js, err := json.Marshal(n)
	return string(js), err
}

// NodeFromXML recreates Node from XML
func NodeFromXML(x string) (n *Node, err error) {
	n = NewNode("")
	err = xml.Unmarshal([]byte(x), n)
	return n, err
}

// ToXML transforms the content of this Engine to XML.
func (n *Node) ToXML() (out string, err error) {
	x, err := xml.Marshal(n)
	return string(x), err
}

// MarshalXML marshalls the node to XML
func (n *Node) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	tokens := []xml.Token{start}

	tokens = append(tokens,
		xml.StartElement{Name: xml.Name{Space: "", Local: "ID"}},
		xml.CharData(n.ID.String()),
		xml.EndElement{Name: xml.Name{Space: "", Local: "ID"}})

	tokens = append(tokens,
		xml.StartElement{Name: xml.Name{Space: "", Local: "Created"}},
		xml.CharData(n.Created.Format(C.TimeFormat)),
		xml.EndElement{Name: xml.Name{Space: "", Local: "Created"}})

	tokens = append(tokens,
		xml.StartElement{Name: xml.Name{Space: "", Local: "Modified"}},
		xml.CharData(n.Modified.Format(C.TimeFormat)),
		xml.EndElement{Name: xml.Name{Space: "", Local: "Modified"}})

	tokens = append(tokens,
		xml.StartElement{Name: xml.Name{Space: "", Local: "Name"}},
		xml.CharData(n.Name),
		xml.EndElement{Name: xml.Name{Space: "", Local: "Name"}})

	for key, value := range n.Data {
		t := xml.StartElement{Name: xml.Name{Space: "", Local: key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{Name: t.Name})
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, t := range tokens {
		err = e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err = e.Flush()
	return err
}

// Serialize this Node.
func (n *Node) Serialize() (out string, err error) {
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

// DeserializeNode a Node.
func DeserializeNode(raw string) (n *Node, err error) {
	switch strings.ToUpper(C.OutputFormat) {
	case "YAML":
		return NodeFromYaml(raw)
	case "JSON":
		return NodeFromJSON(raw)
	case "XML":
		return NodeFromXML(raw)
	default:
		return nil, errors.New("Output format " + C.OutputFormat + " not recognized.")
	}
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
