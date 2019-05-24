package utilities

import (
  "net/http"
)

// Function to set CORS to header response
func SetCORS(w http.ResponseWriter) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
