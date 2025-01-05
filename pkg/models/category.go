package models

import "database/sql"

type Category struct {
	Id   int64          `json:"id"`
	Name string         `json:"Name"`
	Img  sql.NullString `json:"Img"`
	Slug string         `json:"Slug"`
}
