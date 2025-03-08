package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/dto"
	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/middlewares"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
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
func (h Handler) GetRecipesByCategory(w http.ResponseWriter, r *http.Request) {
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
	var category models.Category
	err := h.DB.Get(&category, queries.GetCategoryBySlug, categorySlug)
	if helpers.HandleSQLError(err, "Failed to get category") {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := h.DB.Queryx(queries.GetRecipesWithRelationsByCategory,
		&category.Id, orderBy, limit, offset)
	if err != nil {
		log.Println("Error querying recipes ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var recipes []models.Recipe

	for rows.Next() {
		var recipe models.Recipe
		err = rows.StructScan(&recipe)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, recipe)
	}
	response := dto.CategoryRecipesResponse{
		Category: category,
		Recipes:  recipes,
	}
	helpers.RespondJSON(w, http.StatusOK, response)
}
