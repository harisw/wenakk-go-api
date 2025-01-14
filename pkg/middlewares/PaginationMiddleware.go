package middlewares

import (
	"context"
	"net/http"
	"strconv"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func PaginationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		page := 1
		limit := 30

		if p := r.URL.Query().Get("page"); p != "" {
			if pInt, err := strconv.Atoi(p); err == nil && pInt > 0 {
				page = pInt
			}
		}

		if l := r.URL.Query().Get("limit"); l != "" {
			if lInt, err := strconv.Atoi(l); err == nil && lInt > 0 {
				limit = lInt
			}
		}

		pagination := Pagination{
			Page:   page,
			Limit:  limit,
			Offset: page * limit,
		}
		ctx := context.WithValue(r.Context(), "pagination", pagination)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
