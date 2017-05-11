package kraken

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// This pointer is holding the graph that is currently loaded to RAM.
var current *Graph

func graph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	switch r.Method {
	case "GET":
		g, err := LoadFromDisk(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Not found")
			return
		}
		current = g
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, g.ID.String())
	case "PUT":
		current = NewGraph(name)
		current.SaveToDisk()
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, current.ID.String())
	case "PATCH":
		if current == nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Not found")
			return
		}
		w.WriteHeader(http.StatusOK)
		current.SaveToDisk()
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//Start the service.
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{name}", graph)

	log.Fatal(http.ListenAndServe(Host+":"+strconv.Itoa(Port), router))
}
