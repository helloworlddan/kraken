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

func engine(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	switch r.Method {
	case "GET":
		y, err := E.ToYaml()
		if err != nil {
			respond(w, http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK)
		io.WriteString(w, y)
	default:
		respond(w, http.StatusMethodNotAllowed)
		return
	}
}

func graph(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	vars := mux.Vars(r)
	name := vars["graph"]

	switch r.Method {
	case "GET":
		g, err := E.FindGraph(name)
		if err != nil {
			respond(w, http.StatusNotFound)
			return
		}
		y, err := g.ToYaml()
		if err != nil {
			respond(w, http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK)
		io.WriteString(w, y)
	default:
		respond(w, http.StatusMethodNotAllowed)
		return
	}
}

func node(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	vars := mux.Vars(r)
	name := vars["graph"]
	id := vars["id"]

	g, err := E.FindGraph(name)
	if err != nil {
		respond(w, http.StatusNotFound)
		return
	}
	n, err := g.GetNode(id)
	if err != nil {
		respond(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		y, err := n.ToYaml()
		if err != nil {
			respond(w, http.StatusInternalServerError)
			return
		}
		respond(w, http.StatusOK)
		io.WriteString(w, y)
	default:
		respond(w, http.StatusMethodNotAllowed)
		return
	}
}

func respond(writer http.ResponseWriter, status int) {
	writer.WriteHeader(status)
	logResponse(status)
}

func logRequest(req *http.Request) {
	log.Println("Request: " + req.Method + " => " + req.RequestURI + " from " + req.RemoteAddr)
}

func logResponse(status int) {
	log.Println("Response: " + strconv.Itoa(status))
}

//Start the service.
func Start() {
	log.Println("Starting " + ApplicationName + " Version " + ApplicationVersion)
	E = NewEngine()

	E.LoadDirectory(DefaultStore)
	log.Println("Loaded " + strconv.Itoa(E.CountGraphs()) + " graph(s).")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", engine)
	router.HandleFunc("/{graph}/", graph)
	router.HandleFunc("/{graph}/{id}", node)

	log.Fatal(http.ListenAndServe(Host+":"+strconv.Itoa(Port), router))
}
