package queries

const (
	GetAllOrigins   = `SELECT * from origins;`
	GetOriginBySlug = `SELECT * from origins WHERE slug = $1;`
)
