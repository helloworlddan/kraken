package kraken

import (
	"io"
	"net/http"
)

// ServeEngine hold calls to the Engine type
func ServeEngine(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	switch r.Method {
	case "GET":
		y, err := E.ToYaml()
		if err != nil {
			Respond(w, http.StatusInternalServerError)
			return
		}
		Respond(w, http.StatusOK)
		io.WriteString(w, y)
	default:
		Respond(w, http.StatusMethodNotAllowed)
		return
	}
}
