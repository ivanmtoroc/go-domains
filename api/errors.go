package api

import (
	"net/http"
)

type ErrResponse struct {
	HttpStatusCode  int     `json:"http_status_code"`
	StatusText      string  `json:"status_text"`
}

func (er *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

var ERR404 = &ErrResponse{404, "Resource not found."}
