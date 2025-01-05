package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/harisw/wenakkGoApi/pkg/models"
	"github.com/harisw/wenakkGoApi/pkg/queries"
)

// GetAllRecipes godoc
// @Summary Get all recipes
// @Description Get all recipes
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {array} models.Recipe
// @Router /recipes [get]
func (h handler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "30"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Println("failed to convert page to int", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Println("failed to convert limit to int", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	offset := page * limit
	results, err := h.DB.Query(queries.GetRecipesWithRelations, limit, offset)
	if err != nil {
		log.Println("Error querying recipes ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var recipes = make([]models.Recipe, 0)
	for results.Next() {
		var recipe models.Recipe
		err = results.Scan(&recipe.Id, &recipe.Category.Name, &recipe.Category.Slug,
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
		recipes = append(recipes, recipe)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)
}
