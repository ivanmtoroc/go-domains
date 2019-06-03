package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ivanmtoroc/go-domains/models"
	"github.com/ivanmtoroc/go-domains/responses"
	"github.com/ivanmtoroc/go-domains/utilities"
)

// GetItems is http handler to '/items' endpoint
func GetItems(w http.ResponseWriter, r *http.Request) {
	var offset, limit string
	// Add CORS to response
	utilities.SetCORS(w)

	// Get all url params
	params := r.URL.Query()

	if len(params["offset"]) > 0 {
		offset = string(params["offset"][0])
	}
	if len(params["limit"]) > 0 {
		limit = string(params["limit"][0])
	}

	fmt.Println("consulting valid domains into database")
	// Get domains history
	domains, err := models.GetDomainsDB(offset, limit)
	if err != nil {
		render.Render(w, r, responses.Error500)
		return
	}
	count, errCount := models.GetDomainsCountDB()
	if errCount != nil {
		render.Render(w, r, responses.Error500)
		return
	}
	fmt.Printf("- domains count: %d\n", count)

	// Create and render items response
	if err := render.Render(w, r, responses.CreateItemsResponse(domains, count)); err != nil {
		render.Render(w, r, responses.Error500)
		return
	}
}
