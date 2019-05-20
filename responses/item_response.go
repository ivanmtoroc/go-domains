package responses

import (
  "net/http"
  "go-domains/models"
)

type ItemInfoResponse struct {
  IsDown    bool    `json:"is_down"`
  Logo      string  `json:"logo"`
  Title     string  `json:"title"`
}

type ItemResponse struct {
  Name  string             `json:"name"`
  Info  *ItemInfoResponse  `json:"info"`
}

type ItemsResponse struct {
  Items  []*ItemResponse  `json:"items"`
}

func CreateItemsResponse(domains []*models.Domain) *ItemsResponse {
  var items_response []*ItemResponse = make([]*ItemResponse, 0)

  for _, domain := range domains {
    items_response = append(items_response, createItemResponse(domain))
  }

	return &ItemsResponse{items_response}
}

func createItemResponse(domain *models.Domain) *ItemResponse {
  item_info := &ItemInfoResponse{
    domain.IsDown,
    domain.Logo,
    domain.Title,
  }

  return &ItemResponse{domain.Name, item_info}
}


func (iir *ItemInfoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ir *ItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ir *ItemsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
