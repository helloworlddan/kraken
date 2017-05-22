package kraken

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

// ServeGraph hold calls to the Graph type
func ServeGraph(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	vars := mux.Vars(r)
	name := vars["graph"]
	uid, err := uuid.FromString(name)
	if err != nil {
		g, er := E.FindGraph(name)
		if er != nil {
			Respond(w, http.StatusNotFound)
			return
		}
		uid = g.ID
	}
	g, err := E.GetGraph(uid.String())
	if err != nil {
		Respond(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		y, err := g.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	case "POST":
		n := NewNode("")
		g.AddNode(n)
		y, err := n.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	case "DELETE":
		E.DeleteFromDisk(g)
		E.DropGraph(g)
		Respond(w, http.StatusOK)
		return
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
