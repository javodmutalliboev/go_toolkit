package middleware

import (
	"net/http"
	"os"
	"strings"
)

func Cors(next http.Handler) http.Handler {
	allowedOrigins := []string{os.Getenv("ADMIN_ORIGIN_LOCAL"), os.Getenv("ADMIN_ORIGIN_REMOTE")} // replace with your allowed origins

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Origin header from the request
		origin := r.Header.Get("Origin")

		// Check if the Origin is in the list of allowed origins
		for _, allowedOrigin := range allowedOrigins {
			if strings.EqualFold(origin, allowedOrigin) {
				// Set headers
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		// If it's a preflight request, respond immediately
		if r.Method == "OPTIONS" {
			return
		}

		// Otherwise, pass the request to the next middleware in the chain
		next.ServeHTTP(w, r)
	})
}
