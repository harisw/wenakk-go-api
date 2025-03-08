package queries

const (
	GetRecipesWithRelations = `
		SELECT c.name, c.slug, r.id, r.author_name, r.name, r.total_time, r.date_published, r.description, r.
		       rating, r.calories, r.fat, r.cholesterol, r.sodium, r.carbohydrate, r.fiber, r.
		       sugar, r.protein, r.portions, r.tags, r.instructions, r.images, r.
		       ingredient_qty, r.ingredients, r.external_id, r.slug
		FROM recipes r
		join categories c ON r.category_id = c.id
		LIMIT $1 OFFSET $2;
	`
	GetRecipeWithRelations = `SELECT c.name, c.slug, r.id, r.author_name, r.name, r.total_time, r.date_published, r.description, r.
		       rating, r.calories, r.fat, r.cholesterol, r.sodium, r.carbohydrate, r.fiber, r.
		       sugar, r.pr			log.Println(message+": Query error:", err)
otein, r.portions, r.tags, r.instructions, r.images, r.
		       ingredient_qty, r.ingredients, r.external_id, r.slug
		FROM recipes r
		join categories c ON r.category_id = c.id
		WHERE r.slug = $1;`

	GetRecipesWithRelationsByCategory = `SELECT c.name, c.slug, r.id, r.author_name, r.name, r.total_time, r.date_published, r.description, r.
		       rating, r.calories, r.fat, r.cholesterol, r.sodium, r.carbohydrate, r.fiber, r.
		       sugar, r.protein, r.portions, r.tags, r.instructions, r.images, r.
		       ingredient_qty, r.ingredients, r.external_id, r.slug
		 FROM recipes r
		 join categories c ON r.category_id = c.id
		 WHERE r.category_id=$1
		 ORDER BY $2
		 LIMIT $3 OFFSET $4;`
)
