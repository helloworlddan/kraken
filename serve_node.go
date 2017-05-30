package kraken

import (
	"io"
	"log"
	"net/http"
)

// ServeNode hold calls to the Node type
func ServeNode(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	currentGraph, currentNode, status, err := GetNodeURL(r)
	if err != nil {
		Respond(w, status)
		log.Println(err)
		return
	}

	switch r.Method {
	case "GET":
		y, err := currentNode.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
		return
	case "DELETE":
		currentGraph.DeleteNode(currentNode)
		Respond(w, http.StatusOK)
		return
	case "PATCH":
		update, status, err := GetNodeBody(r.Body)
		if err != nil {
			Respond(w, status)
			log.Println(err)
			return
		}
		currentNode.Update(update)
		out, err := currentNode.Serialize()
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
