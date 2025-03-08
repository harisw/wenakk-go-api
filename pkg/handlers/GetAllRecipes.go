package handlers

import (
	"log"
	"net/http"

	"github.com/harisw/wenakkGoApi/pkg/helpers"
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
func (h Handler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {

	pagination, ok := r.Context().Value("pagination").(middlewares.Pagination)
	if !ok {
		log.Println("Failed to retrieve pagination from middleware")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	limit := pagination.Limit
	offset := pagination.Offset

	rows, err := h.DB.Queryx(queries.GetRecipesWithRelations, limit, offset)
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
	helpers.RespondJSON(w, http.StatusOK, recipes)
}
