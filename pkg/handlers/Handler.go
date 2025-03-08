package handlers

import "github.com/jmoiron/sqlx"

type Handler struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) Handler {
	return Handler{db}
}
