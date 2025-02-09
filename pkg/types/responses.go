package types

import "github.com/harisw/wenakkGoApi/pkg/models"

type OriginRecipesResponse struct {
	Origin  models.Origin   `json:"origin"`
	Recipes []models.Recipe `json:"recipes"`
}

type CategoryRecipesResponse struct {
	Category models.Category `json:"category"`
	Recipes  []models.Recipe `json:"recipes"`
}
