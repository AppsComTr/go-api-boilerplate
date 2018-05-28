package middleware

import (
	"net/http"
)

func AuthMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		inner.ServeHTTP(w, r)
	})
}
