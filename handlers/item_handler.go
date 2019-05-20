package handlers

import (
	"net/http"
	"go-domains/responses"
	"go-domains/models"
	"go-domains/utilities"	
	"github.com/go-chi/render"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	utilities.SetCORS(w)

	domains := models.GetDomainsDB()

	if err := render.Render(w, r, responses.CreateItemsResponse(domains)); err != nil {
		render.Render(w, r, responses.ERROR_500)
		return
	}
}
