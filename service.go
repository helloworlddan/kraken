package kraken

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// E Globally running Engine.
var E *Engine

func graph(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["graph"]

	switch r.Method {
	case "GET":
		g, err := E.FindGraph(name)
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

func engine(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		y, err := E.ToYaml()
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

	E = NewEngine()
	E.LoadDirectory(DefaultStore)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", engine)
	router.HandleFunc("/{graph}/", graph)

	log.Fatal(http.ListenAndServe(Host+":"+strconv.Itoa(Port), router))
}
