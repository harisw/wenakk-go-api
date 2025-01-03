package queries

const GetRecipesWithRelations string = `
		SELECT r.id, c.name, c.slug, o.name, o.slug, r.name, r.total_time, r.date_published, r.description, r.images,
		r.keywords, r.rating, r.calories, r.protein, r.recipe_yield, r.instructions,
		r.recipe_id, r.ingredients
		FROM recipes r
		join categories c ON r.category_id = c.id
		join origins o ON r.origin_id = o.id
		LIMIT $1 OFFSET $2;
	`
