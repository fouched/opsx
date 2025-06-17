package main

import (
	"github.com/fouched/opsx/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func routes(h *handlers.Handlers) http.Handler {
	r := chi.NewRouter()
	addMiddleware(r)

	// Root and health check
	r.Get("/", h.Home)

	r.Route("/apps", func(r chi.Router) {
		r.Get("/", h.AllApps)
	})
	// r.Route("/env", func(r chi.Router) { ... })

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}

func addMiddleware(r *chi.Mux) {
	// sessions
	r.Use(SessionLoad)

	// CORS
	//r.Use(cors.Handler(cors.Options{
	//	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	//	AllowedOrigins: []string{"https://*", "http://*"},
	//	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//	ExposedHeaders:   []string{"Link"},
	//	AllowCredentials: false,
	//	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//}))

	// Recover from panics, logs the panic, and returns HTTP 500
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
}
