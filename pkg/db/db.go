package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	_ = godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portInt, user, password, dbname)

	db, err := sqlx.Connect("postgres", connInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB")
	return db
}

func CloseConnection(db *sqlx.DB) {
	defer db.Close()
}
