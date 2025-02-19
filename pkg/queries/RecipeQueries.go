package queries

const (
	GetRecipesWithRelations = `
		SELECT r.id, c.name, c.slug, o.name, o.slug, r.name, r.total_time, r.date_published, r.description, r.images,
		r.keywords, r.rating, r.calories, r.protein, r.recipe_yield, r.instructions,
		r.recipe_id, r.ingredients
		FROM recipes r
		join categories c ON r.category_id = c.id
		join origins o ON r.origin_id = o.id
		LIMIT $1 OFFSET $2;
	`
	GetRecipeWithRelations = `SELECT r.id, c.name, c.slug, o.name, o.slug, r.name,
	 r.total_time, r.date_published, r.description, r.images,
		r.keywords, r.rating, r.calories, r.protein, r.recipe_yield, r.instructions,
		r.recipe_id, r.ingredients
		FROM recipes r
		join categories c ON r.category_id = c.id
		join origins o on r.origin_id = o.id
		WHERE r.id = $1;`

	GetRecipesWithRelationsByOrigin = `SELECT r.id, c.name, c.slug, o.name, o.slug, r.name,
	 r.total_time, r.date_published, r.description, r.images,
		r.keywords, r.rating, r.calories, r.protein, r.recipe_yield, r.instructions,
		r.recipe_id, r.ingredients
		FROM recipes r
		join categories c ON r.category_id = c.id
		join origins o on r.origin_id = o.id
		WHERE r.origin_id=$1
		LIMIT $2 OFFSET $3;`

	GetRecipesWithRelationsByCategory = `SELECT r.id as id, c.name, c.slug, o.name, o.slug, r.name as name,
		r.total_time, r.date_published, r.description, r.images,
		 r.keywords, r.rating, r.calories, r.protein, r.recipe_yield, r.instructions,
		 r.recipe_id, r.ingredients
		 FROM recipes r
		 join categories c ON r.category_id = c.id
		 join origins o on r.origin_id = o.id
		 WHERE r.category_id=$1
		 ORDER BY $2
		 LIMIT $3 OFFSET $4;`
)
