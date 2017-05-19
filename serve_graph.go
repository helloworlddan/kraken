package kraken

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// ServeGraph hold calls to the Graph type
func ServeGraph(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	vars := mux.Vars(r)
	name := vars["graph"]

	switch r.Method {
	case "GET":
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

		y, err := g.ToYaml()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
	case "POST":
		g := NewGraph(name)
		E.AddGraph(g)
		y, err := g.ToYaml()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
	case "DELETE":
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
		E.DeleteFromDisk(g)
		E.DropGraph(g)
		g = nil
		Respond(w, http.StatusOK)
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
