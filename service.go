package kraken

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var current *Graph

func graph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if r.Method == "GET" {
		g, err := LoadFromDisk(name)
		if err != nil {
			// TODO: codes?
			io.WriteString(w, "Not found")
			return
		}
		current = g
		io.WriteString(w, g.ID.String())
	}

	if r.Method == "PUT" {
		if current == nil {
			// TODO: codes?
			io.WriteString(w, "Not found")
			return
		}
		current.SaveToDisk()
	}
}

//Start the service.
func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{name}", graph)

	log.Fatal(http.ListenAndServe(Host+":"+strconv.Itoa(Port), router))
}
