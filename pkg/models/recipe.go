package models

import (
	"time"

	"github.com/harisw/wenakkGoApi/pkg/types"
)

type Recipe struct {
	Id            int64            `db:"id" json:"id"`
	Category      Category         `db:"category" json:"category"`
	AuthorName    types.NullString `db:"author_name" json:"author_name"`
	Name          string           `db:"name" json:"name"`
	TotalTime     types.NullString `db:"total_time" json:"total_time"`
	DatePublished time.Time        `db:"date_published" json:"date_published"`
	Description   types.NullString `db:"description" json:"description"`
	Rating        types.NullFloat  `db:"rating" json:"rating"`
	Calories      types.NullFloat  `db:"calories" json:"calories"`
	Fat           types.NullFloat  `db:"fat" json:"fat"`
	Cholesterol   types.NullFloat  `db:"cholesterol" json:"cholesterol"`
	Sodium        types.NullFloat  `db:"sodium" json:"sodium"`
	Carbohydrate  types.NullFloat  `db:"carbohydrate" json:"carbohydrate"`
	Fiber         types.NullFloat  `db:"fiber" json:"fiber"`
	Sugar         types.NullFloat  `db:"sugar" json:"sugar"`
	Protein       types.NullFloat  `db:"protein" json:"protein"`
	Portions      types.NullString `db:"portions" json:"portions"`
	Tags          types.JSONArr    `db:"tags" json:"tags" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Instructions  types.JSONArr    `db:"instructions" json:"instructions" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Images        types.JSONArr    `db:"images" json:"images" swaggertype:"string" example:"{\"key\":\"value\"}"`
	IngredientQty types.JSONArr    `db:"ingredient_qty" json:"ingredient_qty" swaggertype:"string" example:"{\"key\":\"value\"}"`
	Ingredients   types.JSONArr    `db:"ingredients" json:"ingredients" swaggertype:"string" example:"{\"key\":\"value\"}"`
	ExternalId    int64            `db:"external_id" json:"external_id"`
	Slug          string           `db:"slug" json:"slug"`
}
