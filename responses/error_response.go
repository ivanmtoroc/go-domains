package responses

import (
	"net/http"
)

type ErrorResponse struct {
	HttpStatusCode  int     `json:"http_status_code"`
	StatusText      string  `json:"status_text"`
}

func (er *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

var ERROR_INVALID_DOMAIN = &ErrorResponse{200, "Invalid domain"}

var ERROR_400 = &ErrorResponse{400, "Bad request"}
var ERROR_404 = &ErrorResponse{404, "Resource not found"}

var ERROR_500 = &ErrorResponse{500, "Internal server error"}
