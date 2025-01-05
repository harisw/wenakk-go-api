package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/db"
	"github.com/harisw/wenakkGoApi/pkg/handlers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Wenakk!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest(DB *sql.DB) {
	h := handlers.New(DB)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/origins", h.GetAllOrigins).Methods("GET")
	myRouter.HandleFunc("/origins/{slug}", h.GetOrigin).Methods("GET")
	myRouter.HandleFunc("/categories", h.GetAllCategories).Methods("GET")
	myRouter.HandleFunc("/categories/{slug}", h.GetCategory).Methods("GET")
	myRouter.HandleFunc("/recipes", h.GetAllRecipes).Methods("GET")
	myRouter.HandleFunc("/recipes/{recipeId}", h.GetRecipe).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	DB := db.Connect()
	handleRequest(DB)
	db.CloseConnection(DB)
}
