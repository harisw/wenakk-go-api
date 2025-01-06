package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harisw/wenakkGoApi/pkg/middlewares"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetAllRecipes godoc
// @Summary Get all recipes
// @Description Get all recipes
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {array} models.Recipe
// @Router /recipes [get]
func (h handler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {

	pagination, ok := r.Context().Value("pagination").(middlewares.Pagination)
	if !ok {
		log.Println("Failed to retrieve pagination from middleware")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	page := pagination.Page
	limit := pagination.Limit
	offset := page * limit
	results, err := h.DB.Query(queries.GetRecipesWithRelations, limit, offset)
	if err != nil {
		log.Println("Error querying recipes ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var recipes = make([]models.Recipe, 0)
	for results.Next() {
		var recipe models.Recipe
		err = results.Scan(&recipe.Id, &recipe.Category.Name, &recipe.Category.Slug,
			&recipe.Origin.Name, &recipe.Origin.Slug,
			&recipe.Name, &recipe.TotalTime, &recipe.DatePublished, &recipe.Description,
			&recipe.Images, &recipe.Keywords, &recipe.Rating, &recipe.Calories,
			&recipe.Protein, &recipe.RecipeYield, &recipe.Instructions, &recipe.RecipeId,
			&recipe.Ingredients)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, recipe)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)
}
