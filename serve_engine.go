package kraken

import (
	"io"
	"log"
	"net/http"
)

// ServeEngine hold calls to the Engine type
func ServeEngine(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	switch r.Method {
	case "GET":
		y, err := E.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	case "POST":
		g := NewGraph("")
		E.AddGraph(g)
		y, err := g.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
