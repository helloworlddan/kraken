package kraken

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var e *Engine

func graph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["graph"]

	switch r.Method {
	case "GET":
		g, err := e.FindGraph(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		y, err := g.ToYaml()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, y)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//Start the service.
func Start() {

	e = NewEngine()

	// TODO: Load all found graphs
	g, err := LoadFromDisk("Flights")
	if err != nil {
		log.Println(err)
	}
	e.AddGraph(g)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{graph}/", graph)

	log.Fatal(http.ListenAndServe(Host+":"+strconv.Itoa(Port), router))
}
