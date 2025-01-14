package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/middlewares"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

type Response struct {
	Origin  models.Origin   `json:"origin"`
	Recipes []models.Recipe `json:"recipes"`
}

// GetRecipesByOrigin godoc
// @Summary Get all recipes by Origins
// @Description Get all recipes paginated by its origin
// @Tags recipes
// @Accept json
// @Produce json
// @Param slug path string true "Origin Slug"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} Response
// @Router /recipes/origin/{slug} [get]
func (h handler) GetRecipesByOrigin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	originSlug := vars["slug"]
	pagination, ok := r.Context().Value("pagination").(middlewares.Pagination)
	if !ok {
		log.Println("Failed to retrieve pagination from middleware")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	limit := pagination.Limit
	offset := pagination.Offset

	result, err := h.DB.Query(queries.GetOriginBySlug, originSlug)
	if err != nil {
		log.Println("Error getting origin")
	}
	var origin models.Origin
	for result.Next() {
		err = result.Scan(&origin.Id, &origin.Name, &origin.Slug)
		if err != nil {
			log.Println("failed to scan origin ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	results, err := h.DB.Query(queries.GetRecipesWithRelationsByOrigin, &origin.Id,
		limit, offset)
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
	response := Response{
		Origin:  origin,
		Recipes: recipes,
	}
	helpers.RespondJSON(w, http.StatusOK, response)
}
