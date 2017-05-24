package kraken

import (
	"io"
	"io/ioutil"
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
	case "PUT":
		// Should not be used
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, string(body))
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
