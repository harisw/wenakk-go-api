package models

type Category struct {
	Id    int64   `json:"id"`
	Name  string  `json:"Name"`
	Img   *string `json:"Img"`
	Slug  string  `json:"Slug"`
	Count int64   `json:"Count"`
}
