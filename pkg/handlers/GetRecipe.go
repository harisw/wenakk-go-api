package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetRecipe godoc
// @Summary Get recipe by id
// @Description Get recipe by id
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param recipeId path string true "Recipe ID"
// @Success 200 {object} models.Recipe
// @Router /recipes/{recipeId} [get]
func (h Handler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	var recipe models.Recipe
	err := h.DB.Get(&recipe, queries.GetRecipeWithRelations, slug)
	if helpers.HandleSQLError(err, "Failed to get recipe") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.RespondJSON(w, http.StatusOK, recipe)
}
