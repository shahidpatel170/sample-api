package main

import (
	c "context"
	"net/http"
	"strings"
)

func tokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientId = r.PathValue("clientId")
		clientProfile, ok := database[clientId]
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		token := r.Header.Get("Authorization")
		if !isValidToken(clientProfile, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := c.WithValue(r.Context(), "clientProfile", clientProfile)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}

func isValidToken(cprof ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == cprof.Token
	}
	return false
}
