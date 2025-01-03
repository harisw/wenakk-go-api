package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

func (h handler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeId := vars["recipeId"]

	result, err := h.DB.Query(queries.GetRecipeWithRelations, recipeId)
	if err != nil {
		log.Println("Error querying recipe ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var recipe models.Recipe
	for result.Next() {
		err = result.Scan(&recipe.Id, &recipe.Category.Name, &recipe.Category.Slug,
			&recipe.Origin.Name, &recipe.Origin.Slug,
			&recipe.Name, &recipe.TotalTime, &recipe.DatePublished, &recipe.Description,
			&recipe.Images, &recipe.Keywords, &recipe.Rating, &recipe.Calories,
			&recipe.Protein, &recipe.RecipeYield, &recipe.Instructions, &recipe.RecipeId,
			&recipe.Ingredients)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe)
}
