package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harisw/wenakkGoApi/pkg/models"
)

func (h handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query("SELECT * FROM categories;")
	if err != nil {
		log.Println("Error querying categories ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var categories = make([]models.Category, 0)
	for results.Next() {
		var category models.Category
		err = results.Scan(&category.Id, &category.Name, &category.Slug)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}
