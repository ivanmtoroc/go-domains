package responses

import (
	"net/http"
)

// ErrorResponse structure to create errors JSON responses
type ErrorResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

// Render is a method of Renderer interface of go-chi/render for managing response payloads
func (er *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Error400 is error response to a bad request
var Error400 = &ErrorResponse{
	Status:     400,
	StatusText: "Bad request",
}

// Error404 is error response to handle a invalid domain name request
var Error404 = &ErrorResponse{
	Status:     404,
	StatusText: "Invalid domain name",
}

// Error500 is error respose to a internar server error
var Error500 = &ErrorResponse{
	Status:     500,
	StatusText: "Internal server error",
}
