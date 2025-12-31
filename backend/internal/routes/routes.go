package routes

import (
	"encoding/json"
	"net/http"
)

// Router
type Router struct{}

type Handler interface {
	ServHTTP(w http.ResponseWriter, r *http.Request)
}

// NewRouter create a new instance
func NewRouter() *Router {
	return &Router{}
}

// ServeHTTP implements http.Handler.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	// manual routing - check path and method

	switch {
	case method == http.MethodGet && path == "/health":
		router.health(w, r)

	// Add more routes as we grow

	default:
		// No route matched - return 404
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// health handles GET /health
func (router *Router) health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	json.NewEncoder(w).Encode(response)
}
