package models

import (
	"encoding/json"
	"time"
)

type Recipe struct {
	Id            int64           `json:"Id"`
	Category      Category        `json:"Category"`
	Origin        Origin          `json:"Origin"`
	Name          string          `json:"Name"`
	TotalTime     string          `json:"Total_time"`
	DatePublished time.Time       `json:"Date_published"`
	Description   string          `json:"Description"`
	Images        json.RawMessage `json:"Images" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Keywords      json.RawMessage `json:"Keywords" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Rating        float32         `json:"Rating"`
	Calories      float32         `json:"Calories"`
	Protein       float32         `json:"Protein"`
	RecipeYield   *string         `json:"Recipe_yield"`
	Instructions  json.RawMessage `json:"Instructions" swaggertype:"string" example:"{\"key\":\"value\"}"`
	RecipeId      int64           `json:"Recipe_id"`
	Ingredients   *string         `json:"Ingredients"`
}
