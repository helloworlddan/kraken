package core

import (
	"io"
	"log"
	"net/http"
)

// ServeGraph hold calls to the Graph type
func ServeGraph(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	current, status, err := GetGraphURL(r)
	if err != nil {
		Respond(w, status)
		log.Println(err)
		return
	}

	switch r.Method {
	case "GET": // Get an existing Graph.
		y, err := current.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	case "DELETE": // Delete an existing Graph.
		E.DropGraph(current)
		current = nil
		Respond(w, http.StatusOK)
		return
	case "PATCH": // Update an existing Graph with specified body.
		update, status, err := GetGraphBody(r.Body)
		if err != nil {
			Respond(w, status)
			log.Println(err)
			return
		}
		current.Update(update)
		out, err := current.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "POST": // Post empty Node into Graph.
		n := NewNode("")
		current.AddNode(n)
		y, err := n.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	case "PUT": // Create a new Node with the specified body.
		update, status, err := GetNodeBody(r.Body)
		if err != nil {
			Respond(w, status)
			log.Println(err)
			return
		}
		n := NewNode("")
		n.Update(update)
		current.AddNode(n)
		out, err := n.Serialize()
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
