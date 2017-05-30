package kraken

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Respond to a web request
func Respond(writer http.ResponseWriter, status int) {
	switch strings.ToUpper(C.OutputFormat) {
	case "YAML":
		writer.Header().Set("Content-Type", "application/yaml")
	case "JSON":
		writer.Header().Set("Content-Type", "application/json")
	case "XML":
		writer.Header().Set("Content-Type", "application/xml")
	default:
		writer.Header().Set("Content-Type", "text/plain")
	}
	writer.WriteHeader(status)
	LogResponse(status)
}

// LogRequest is logging incoming requests.
func LogRequest(req *http.Request) {
	log.Println("Request: " + req.Method + " => " + req.RequestURI + " from " + req.RemoteAddr)
}

// LogResponse is logging outgoing responses.
func LogResponse(status int) {
	log.Println("Response: " + strconv.Itoa(status))
}

// GetGraphBody delivers a Graph from a Body context.
func GetGraphBody(body io.ReadCloser) (update *Graph, status int, err error) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	update, err = DeserializeGraph(string(content))
	if err != nil {
		return update, http.StatusBadRequest, err
	}
	return update, http.StatusOK, nil
}

// GetGraphURL delivers a graph from an URL context.
func GetGraphURL(req *http.Request) (current *Graph, status int, err error) {
	vars := mux.Vars(req)
	name := vars["graph"]
	uid, err := uuid.FromString(name)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	current, err = E.GetGraph(uid.String())
	if err != nil {
		return current, http.StatusNotFound, err
	}
	return current, http.StatusOK, nil
}
