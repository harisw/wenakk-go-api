package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetCategory godoc
// @Summary Get category by slug
// @Description Get category by slug
// @Tags categories
// @Accept  json
// @Produce  json
// @Param slug path string true "Category Slug"
// @Success 200 {object} models.Category
// @Router /categories/{slug} [get]
func (h Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	var category models.Category
	err := h.DB.Get(&category, queries.GetCategoryBySlug, slug)

	if helpers.HandleSQLError(err, "Failed to get category") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.RespondJSON(w, http.StatusOK, category)
}
