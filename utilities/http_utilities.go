package utilities

import (
	"net/http"
)

// SetCORS set CORS policies to header response to allow all origins
func SetCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
