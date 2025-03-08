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
func (h Handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	rows, err := h.DB.Queryx(queries.GetAllCategories)
	if err != nil {
		log.Println("Error querying categories ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.StructScan(&category)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	helpers.RespondJSON(w, http.StatusOK, categories)
}
