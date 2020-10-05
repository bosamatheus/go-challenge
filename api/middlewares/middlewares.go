package middlewares

import (
	"errors"
	"mercafacil-challenge/api/auth"
	"mercafacil-challenge/api/responses"
	"net/http"
)

func MiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
