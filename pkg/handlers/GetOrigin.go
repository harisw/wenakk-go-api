package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

func (h handler) GetOrigin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	result, err := h.DB.Query(queries.GetOriginBySlug, slug)
	if err != nil {
		log.Println("Error querying origin ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var origin models.Origin
	for result.Next() {
		err = result.Scan(&origin.Id, &origin.Name, &origin.Slug)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(origin)
}
