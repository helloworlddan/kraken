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
	case "GET": // Get this entire Engine.
		out, err := E.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "POST": // Create a new empty Graph
		current := NewGraph("")
		E.AddGraph(current)
		out, err := current.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "PUT": // Create a new graph with the specified body
		update, status, err := GetGraphBody(r.Body)
		if err != nil {
			Respond(w, status)
			log.Println(err)
			return
		}
		current := NewGraph("")
		current.Update(update)
		E.AddGraph(current)
		out, err := current.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
