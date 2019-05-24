package responses

import (
  "net/http"
)

// Struct used to create error JSON responses
type ErrorResponse struct {
  HttpStatusCode  int     `json:"http_status_code"`
  StatusText      string  `json:"status_text"`
}

// Method used to Renderer interface of go-chi/render for managing response payloads
func (er *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
  return nil
}


var ERROR_400 = &ErrorResponse{400, "Bad request"}
var ERROR_INVALID_DOMAIN = &ErrorResponse{404, "Invalid domain"}

var ERROR_500 = &ErrorResponse{500, "Internal server error"}
