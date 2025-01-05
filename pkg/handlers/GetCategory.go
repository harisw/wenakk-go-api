package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

func (h handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	result, err := h.DB.Query(queries.GetCategoryBySlug, slug)
	if err != nil {
		log.Println("Error querying category ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var category models.Category
	for result.Next() {
		err = result.Scan(&category.Id, &category.Name, &category.Img, &category.Slug)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}
