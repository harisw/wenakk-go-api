package dto

import "github.com/harisw/wenakkGoApi/pkg/models"

type CategoryRecipesResponse struct {
	Category models.Category `json:"category"`
	Recipes  []models.Recipe `json:"recipes"`
}
