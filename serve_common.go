package kraken

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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
