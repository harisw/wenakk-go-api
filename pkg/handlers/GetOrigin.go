package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/helpers"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetOrigin godoc
// @Summary Get origin by slug
// @Description Get origin by slug
// @Tags origins
// @Accept  json
// @Produce  json
// @Param slug path string true "Origin Slug"
// @Success 200 {object} models.Origin
// @Router /origins/{slug} [get]
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
	helpers.RespondJSON(w, http.StatusOK, origin)
}
