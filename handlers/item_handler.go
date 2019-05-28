package handlers

import (
  "fmt"
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

  params := r.URL.Query()

  skip := params["skip"]
  skip_value := ""
  if len(skip) > 0 {
    skip_value = string(skip[0])
  }

  limit := params["limit"]
  limit_value := ""
  if len(limit) > 0 {
    limit_value = string(limit[0])
  }

  fmt.Println("consulting valid domains into database")

  // Get domains history
  domains := models.GetDomainsDB(skip_value, limit_value)
  count := models.GetDomainsCountDB()
  fmt.Printf("- total domains count: %d\n", count)

  // Render items response
  if err := render.Render(w, r, responses.CreateItemsResponse(domains, count)); err != nil {
    render.Render(w, r, responses.ERROR_500)
    return
  }
}
