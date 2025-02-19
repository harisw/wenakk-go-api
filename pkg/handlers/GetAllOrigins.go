package handlers

import (
	"log"
	"net/http"

	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetAllOrigins godoc
// @Summary Get all origins
// @Description Get all origins
// @Tags origins
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Origin
// @Router /origins [get]
func (h handler) GetAllOrigins(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query(queries.GetAllOrigins)
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
	helpers.RespondJSON(w, http.StatusOK, origins)
}
