package responses

import (
	"net/http"

	"github.com/ivanmtoroc/go-domains/models"
)

// ItemInfoResponse structure to create items info JSON responses
type ItemInfoResponse struct {
	IsDown bool   `json:"isDown"`
	Logo   string `json:"logo"`
	Title  string `json:"title"`
}

// ItemResponse structure to create item JSON responses
type ItemResponse struct {
	Name string            `json:"name"`
	Info *ItemInfoResponse `json:"info"`
}

// ItemsResponse structure to create items JSON responses
type ItemsResponse struct {
	Items []*ItemResponse `json:"items"`
	Count int             `json:"totalItems"`
}

// CreateItemsResponse create new items response by domains
func CreateItemsResponse(domains []*models.Domain, count int) *ItemsResponse {
	var itemsResponse = make([]*ItemResponse, 0)
	// Create array of domains responses
	for _, domain := range domains {
		itemsResponse = append(itemsResponse, createItemResponse(domain))
	}
	// Create items respose
	return &ItemsResponse{
		Items: itemsResponse,
		Count: count,
	}
}

// createItemResponse create new item response by domain
func createItemResponse(domain *models.Domain) *ItemResponse {
	// Get info to new item
	itemInfo := &ItemInfoResponse{
		IsDown: domain.IsDown,
		Logo:   domain.Logo,
		Title:  domain.Title,
	}
	// Create item respose
	return &ItemResponse{
		Name: domain.Name,
		Info: itemInfo,
	}
}

// Render is a method of Renderer interface of go-chi/render for managing response payloads
func (ir *ItemsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
