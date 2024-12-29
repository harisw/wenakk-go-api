package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harisw/wenakkGoApi/pkg/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Wenakk!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest(DB *sql.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	DB := db.Connect()
	handleRequest(DB)
	db.CloseConnection(DB)
}
