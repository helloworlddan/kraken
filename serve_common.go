package kraken

import (
	"log"
	"net/http"
	"strconv"
)

// Respond to a web request
func Respond(writer http.ResponseWriter, status int) {
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
