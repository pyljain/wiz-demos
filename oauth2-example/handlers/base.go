package handlers

import (
	"net/http"
)

// New creates a new HTTP server and sets up routing.
func New() http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Root: Serve files from the "templates/" directory
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	// OauthGoogle: Set up route handlers for Google OAuth
	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)

	return mux
}
