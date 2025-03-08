package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/harisw/wenakkGoApi/docs"
	"github.com/harisw/wenakkGoApi/pkg/db"
	"github.com/harisw/wenakkGoApi/pkg/handlers"
	"github.com/harisw/wenakkGoApi/pkg/middlewares"
	"github.com/jmoiron/sqlx"
	httpSwagger "github.com/swaggo/http-swagger"
)

func handleRequest(db *sqlx.DB) {
	h := handlers.New(db)
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.PaginationMiddleware)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/categories", h.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{slug}", h.GetCategory).Methods("GET")
	router.HandleFunc("/recipes", h.GetAllRecipes).Methods("GET")
	router.HandleFunc("/recipes/category/{slug}", h.GetRecipesByCategory).Methods("GET")
	router.HandleFunc("/recipes/{slug}", h.GetRecipe).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

//	@title			Wenakk API
//	@version		1.0
//	@description	This is API Documentation for Wenakk API
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Hari Setiawan
//	@contact.email	haristw16@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {
	DB := db.Connect()
	handleRequest(DB)
	db.CloseConnection(DB)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Wenakk!")
	fmt.Println("Endpoint Hit: homePage")
}
