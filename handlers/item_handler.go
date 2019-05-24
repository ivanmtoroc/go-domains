package handlers

import (
  "net/http"
  "go-domains/responses"
  "go-domains/models"
  "go-domains/utilities"
  "github.com/go-chi/render"
)

// Function to handle requests to '/items' endpoint
func GetItems(w http.ResponseWriter, r *http.Request) {
  // Add CORS to response
  utilities.SetCORS(w)

  // Get domains history
  domains := models.GetDomainsDB()

  // Render items response
  if err := render.Render(w, r, responses.CreateItemsResponse(domains)); err != nil {
    render.Render(w, r, responses.ERROR_500)
    return
  }
}
