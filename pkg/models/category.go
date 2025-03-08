package models

import "github.com/harisw/wenakkGoApi/pkg/types"

type Category struct {
	Id    int64            `db:"id" json:"id"`
	Name  string           `db:"name" json:"name"`
	Img   types.NullString `db:"img" json:"img"`
	Slug  string           `db:"slug" json:"slug"`
	Count int64            `db:"count" json:"count"`
}
