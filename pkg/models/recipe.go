package models

import (
	"time"

	t "github.com/harisw/wenakkGoApi/pkg/types"
)

type Recipe struct {
	Id            int64        `db:"id" json:"id"`
	Category      Category     `db:"category" json:"category"`
	AuthorName    t.NullString `db:"author_name" json:"author_name"`
	Name          string       `db:"name" json:"name"`
	TotalTime     t.NullString `db:"total_time" json:"total_time"`
	DatePublished time.Time    `db:"date_published" json:"date_published"`
	Description   t.NullString `db:"description" json:"description"`
	Rating        t.NullFloat  `db:"rating" json:"rating"`
	Calories      t.NullFloat  `db:"calories" json:"calories"`
	Fat           t.NullFloat  `db:"fat" json:"fat"`
	Cholesterol   t.NullFloat  `db:"cholesterol" json:"cholesterol"`
	Sodium        t.NullFloat  `db:"sodium" json:"sodium"`
	Carbohydrate  t.NullFloat  `db:"carbohydrate" json:"carbohydrate"`
	Fiber         t.NullFloat  `db:"fiber" json:"fiber"`
	Sugar         t.NullFloat  `db:"sugar" json:"sugar"`
	Protein       t.NullFloat  `db:"protein" json:"protein"`
	Portions      t.NullString `db:"portions" json:"portions"`
	Tags          t.JSONArr    `db:"tags" json:"tags" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Instructions  t.JSONArr    `db:"instructions" json:"instructions" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Images        t.JSONArr    `db:"images" json:"images" swaggertype:"string" example:"{\"key\":\"value\"}"`
	IngredientQty t.JSONArr    `db:"ingredient_qty" json:"ingredient_qty" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Ingredients   t.JSONArr    `db:"ingredients" json:"ingredients" swaggertype:"string" example:"{\"key\":\"value\"}"`
	ExternalId    int64        `db:"external_id" json:"external_id"`
	Slug          string       `db:"slug" json:"slug"`
}
