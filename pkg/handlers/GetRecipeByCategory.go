package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/middlewares"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
	"github.com/harisw/wenakkGoApi/pkg/types"
)

// GetRecipesByCategory godoc
// @Summary Get all recipes by Categories
// @Description Get all recipes paginated by its category
// @Tags recipes
// @Accept json
// @Produce json
// @Param slug path string true "Category Slug"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param orderBy query string false "OrderBy"
// @Success 200 {object} types.CategoryRecipesResponse
// @Router /recipes/category/{slug} [get]
func (h handler) GetRecipesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categorySlug := vars["slug"]
	pagination, ok := r.Context().Value("pagination").(middlewares.Pagination)
	if !ok {
		log.Println("Failed to retrieve pagination from middleware")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	limit := pagination.Limit
	offset := pagination.Offset
	orderBy := pagination.OrderBy

	result, err := h.DB.Query(queries.GetCategoryBySlug, categorySlug)
	if err != nil {
		log.Println("Error getting category")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()

	var category models.Category
	for result.Next() {
		err = result.Scan(&category.Id, &category.Name, &category.Img, &category.Slug)
		if err != nil {
			log.Println("failed to scan category ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	results, err := h.DB.Query(queries.GetRecipesWithRelationsByCategory,
		&category.Id, orderBy, limit, offset)
	if err != nil {
		log.Println("Error querying recipes ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer results.Close()
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
	response := types.CategoryRecipesResponse{
		Category: category,
		Recipes:  recipes,
	}
	helpers.RespondJSON(w, http.StatusOK, response)
}
