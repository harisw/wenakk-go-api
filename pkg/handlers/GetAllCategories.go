package handlers

import (
	"log"
	"net/http"

	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query(queries.GetAllCategories)
	if err != nil {
		log.Println("Error querying categories ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var categories = make([]models.Category, 0)
	for results.Next() {
		var category models.Category
		err = results.Scan(&category.Id, &category.Name, &category.Img, &category.Slug)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	helpers.RespondJSON(w, http.StatusOK, categories)
}
