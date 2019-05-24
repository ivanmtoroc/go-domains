package responses

import (
  "net/http"
  "go-domains/models"
)

// Structs used to create items JSON responses
type ItemInfoResponse struct {
  IsDown  bool    `json:"is_down"`
  Logo    string  `json:"logo"`
  Title   string  `json:"title"`
}

type ItemResponse struct {
  Name  string             `json:"name"`
  Info  *ItemInfoResponse  `json:"info"`
}

type ItemsResponse struct {
  Items  []*ItemResponse  `json:"items"`
}

// Create new items response by domains
func CreateItemsResponse(domains []*models.Domain) *ItemsResponse {
  var items_response []*ItemResponse = make([]*ItemResponse, 0)

  // Create array of domains responses
  for _, domain := range domains {
    items_response = append(items_response, createItemResponse(domain))
  }

  return &ItemsResponse{items_response}
}

// Create new item response by domain
func createItemResponse(domain *models.Domain) *ItemResponse {
  // Get info to new item
  item_info := &ItemInfoResponse{
    domain.IsDown,
    domain.Logo,
    domain.Title,
  }

  return &ItemResponse{domain.Name, item_info}
}

// Method used to Renderer interface of go-chi/render for managing response payloads
func (ir *ItemsResponse) Render(w http.ResponseWriter, r *http.Request) error {
  return nil
}
