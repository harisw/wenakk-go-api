package queries

const (
	GetAllCategories  = `SELECT * from categories;`
	GetCategoryBySlug = `SELECT * from categories WHERE slug = $1;`
)
