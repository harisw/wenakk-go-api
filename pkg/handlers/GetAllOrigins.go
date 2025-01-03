package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harisw/wenakkGoApi/pkg/models"
)

func (h handler) GetAllOrigins(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query("SELECT * FROM origins;")
	if err != nil {
		log.Println("Error querying origins ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var origins = make([]models.Origin, 0)
	for results.Next() {
		var origin models.Origin
		err = results.Scan(&origin.Id, &origin.Name, &origin.Slug)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		origins = append(origins, origin)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(origins)
}
