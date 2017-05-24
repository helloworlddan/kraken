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
		out, err := E.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "PATCH": // Update an existing Graph with specified body.
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusInternalServerError)
			return
		}
		update, err := DeserializeGraph(string(body))
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusBadRequest)
			return
		}
		g, err := E.GetGraph(update.ID.String())
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusNotFound)
			return
		}

		g.Update(update)

		out, err := g.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "PUT": // Create a new graph with the specified body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusInternalServerError)
			return
		}
		update, err := DeserializeGraph(string(body))
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusBadRequest)
			return
		}
		g, err := E.GetGraph(update.ID.String())
		if g != nil {
			log.Println(err)
			Respond(w, http.StatusConflict)
			return
		}

		g = NewGraph("")
		g.Update(update)
		E.AddGraph(g)

		out, err := g.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "POST": // Create a new Graph without anything
		g := NewGraph("")
		E.AddGraph(g)
		out, err := g.Serialize()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, out)
		return
	case "DELETE": // Delete an existing Graph.
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusInternalServerError)
			return
		}
		update, err := DeserializeGraph(string(body))
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusBadRequest)
			return
		}
		g, err := E.GetGraph(update.ID.String())
		if err != nil {
			log.Println(err)
			Respond(w, http.StatusNotFound)
			return
		}

		E.DropGraph(g)
		g = nil
		Respond(w, http.StatusOK)
		return
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
