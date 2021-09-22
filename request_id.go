package main

import (
	"context"
	"net/http"
)

type requestIDContextKey struct{}

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(
			w,
			r.WithContext(
				context.WithValue(
					r.Context(),
					requestIDContextKey{},
					"badfood",
				),
			),
		)
	})
}

func GetRequestID(r *http.Request) string {
	v, ok := r.Context().Value(requestIDContextKey{}).(string)
	if !ok {
		return "deadbeef"
	}

	return v
}
