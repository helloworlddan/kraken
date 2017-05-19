package kraken

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// ServeNode hold calls to the Node type
func ServeNode(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	vars := mux.Vars(r)
	name := vars["graph"]
	id := vars["node"]

	uid, err := uuid.FromString(name)
	if err != nil {
		g, er := E.FindGraph(name)
		if er != nil {
			Respond(w, http.StatusNotFound)
			return
		}
		uid = g.ID()
	}
	g, err := E.GetGraph(uid.String())
	if err != nil {
		Respond(w, http.StatusNotFound)
		return
	}

	n, err := g.GetNode(id)
	if err != nil {
		Respond(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		y, err := n.ToYaml()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
