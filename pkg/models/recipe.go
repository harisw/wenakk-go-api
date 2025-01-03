package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Recipe struct {
	Id            int64           `json:"Id"`
	CategoryId    int64           `json:"CategoryId"`
	OriginId      int64           `json:"OriginId"`
	Name          string          `json:"Name"`
	TotalTime     string          `json:"Total_time"`
	DatePublished time.Time       `json:"Date_published"`
	Description   string          `json:"Description"`
	Images        json.RawMessage `json:"Images"`
	Keywords      json.RawMessage `json:"Keywords"`
	Rating        float32         `json:"Rating"`
	Calories      float32         `json:"Calories"`
	Protein       float32         `json:"Protein"`
	RecipeYield   sql.NullString  `json:"Recipe_yield"` // Changed to pointer to accept null value
	Instructions  json.RawMessage `json:"Instructions"`
	RecipeId      int64           `json:"Recipe_id"`
	Ingredients   sql.NullString  `json:"Ingredients"`
}
