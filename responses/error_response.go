package responses

import (
	"net/http"
)

// ErrorResponse structure to create errors JSON responses
type ErrorResponse struct {
	HTTPSStatusCode int    `json:"http_status_code"`
	StatusText      string `json:"status_text"`
}

// Render is a method of Renderer interface of go-chi/render for managing response payloads
func (er *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Error400 is error response to a bad request
var Error400 = &ErrorResponse{
	HTTPSStatusCode: 400,
	StatusText:      "Bad request",
}

// ErrorInvalidDomain is error response to handle a invalid domain name request
var ErrorInvalidDomain = &ErrorResponse{
	HTTPSStatusCode: 404,
	StatusText:      "Invalid domain",
}

// Error500 is error respose to a internar server error
var Error500 = &ErrorResponse{
	HTTPSStatusCode: 500,
	StatusText:      "Internal server error",
}
