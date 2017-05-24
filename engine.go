package kraken

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

// Engine holding all Graphs.
type Engine struct {
	ID      uuid.UUID
	Started time.Time
	Graphs  []*Graph
}

// Inspect inspects this Engine.
func (e *Engine) Inspect() {
	fmt.Printf("ID:\t\t%s\n", e.ID)
	fmt.Printf("Type:\t\tEngine\n")
	fmt.Printf("Started:\t%s\n", e.Started.Format(C.TimeFormat))
	fmt.Printf("Graphs:\t\t%d\n", e.CountGraphs())
	fmt.Printf("\n")
}

// GetGraph tries to find a Graph based on an ID.
func (e *Engine) GetGraph(id string) (*Graph, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	for _, elem := range e.Graphs {
		if elem.ID == uid {
			return elem, nil
		}
	}
	return nil, errors.New("graph not found")
}

// FindGraph tries to find a Graph by its name.
func (e *Engine) FindGraph(name string) (*Graph, error) {
	for _, elem := range e.Graphs {
		if elem.Name == name {
			return elem, nil
		}
	}
	return nil, errors.New("graph not found")
}

// ToYaml transforms the content of this Engine to yaml.
func (e *Engine) ToYaml() (string, error) {
	yam, err := yaml.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(yam), nil
}

// ToJSON transforms the content of this Engine to yaml.
func (e *Engine) ToJSON() (string, error) {
	js, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

// ToXML transforms the content of this Engine to xml.
func (e *Engine) ToXML() (string, error) {
	x, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(x), nil
}

// Serialize this Engine.
func (e *Engine) Serialize() (string, error) {
	switch strings.ToUpper(C.OutputFormat) {
	case "YAML":
		return e.ToYaml()
	case "JSON":
		return e.ToJSON()
	case "XML":
		return e.ToXML()
	default:
		return "", errors.New("Output format " + C.OutputFormat + " not recognized.")
	}
}

// AddGraph to Engine.
func (e *Engine) AddGraph(g *Graph) {
	index := -1
	for i, elem := range e.Graphs {
		if g == elem {
			index = i
		}
	}
	if index == -1 {
		e.Graphs = append(e.Graphs, g)
	}
}

// DropGraph from Engine.
func (e *Engine) DropGraph(g *Graph) {
	index := -1
	for i, elem := range e.Graphs {
		if g == elem {
			index = i
		}
	}
	if index > -1 {
		e.Graphs = append(e.Graphs[:index], e.Graphs[index+1:]...)
	}
}

// CountGraphs counts the total number of all Graphs in this Engine.
func (e *Engine) CountGraphs() int {
	return len(e.Graphs)
}

// LoadDirectory loads all .kraken files in the given directory.
func (e *Engine) LoadDirectory(path string) error {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), C.FileSuffix) {
			name := strings.TrimSuffix(f.Name(), C.FileSuffix)
			g, err := e.ReadFromDisk(name)
			if err != nil {
				return nil
			}
			e.AddGraph(g)
		}
	}
	return nil
}

// CleanupStaleDBFiles deletes stales persistence files.
func (e *Engine) CleanupStaleDBFiles() (int, error) {
	deleted := 0
	var last error
	files, _ := ioutil.ReadDir(C.DefaultStore)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), C.FileSuffix) {
			stale := true
			for _, g := range e.Graphs {
				if (g.ID.String() + C.FileSuffix) == f.Name() {
					stale = false
					break
				}
			}
			if stale {
				err := os.Remove(f.Name())
				if err != nil {
					last = err
				}
				deleted++
			}
		}
	}
	return deleted, last
}

// DeleteFromDisk deletes the database store from disk.
func (e *Engine) DeleteFromDisk(g *Graph) {
	fileName := g.Name + C.FileSuffix
	os.Remove(fileName)
}

// WriteAllToDisk saves all Graphs associated with this Engine.
func (e *Engine) WriteAllToDisk() (int, error) {
	for _, elem := range e.Graphs {
		if err := e.WriteToDisk(elem); err != nil {
			return 0, err
		}
	}
	return len(e.Graphs), nil
}

// WriteToDisk writes the content of this graph to disk.
func (e *Engine) WriteToDisk(g *Graph) error {
	g.Saved = time.Now()
	fileName := g.ID.String() + C.FileSuffix

	y, err := g.ToYaml()
	if err != nil {
		return err
	}
	data := []byte(y)

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadFromDisk loads the graph from the disk.
// Needs the name of the graph to load.
func (e *Engine) ReadFromDisk(id string) (*Graph, error) {
	fileName := id + C.FileSuffix

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	gra, err := GraphFromYaml(string(data))
	if err != nil {
		return nil, err
	}

	return gra, nil
}

// NewEngine creates a brand new Engine.
func NewEngine() *Engine {
	return &Engine{
		ID:      uuid.NewV4(),
		Graphs:  make([]*Graph, 0),
		Started: time.Now(),
	}
}
