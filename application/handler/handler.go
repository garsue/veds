package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garsue/veds/application"
	vmiddleware "github.com/garsue/veds/application/middleware"
	"github.com/garsue/veds/domain/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewHandler returns http handler for application.
func NewHandler(app *application.App) http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", index(app))

	r.With(vmiddleware.ContentTypeJSON).Route("/namespaces", func(r chi.Router) {
		r.Get("/", namespaces(app))
	})

	r.With(vmiddleware.ContentTypeJSON).Route("/kinds", func(r chi.Router) {
		r.Get("/", kinds(app))
	})

	r.With(vmiddleware.ContentTypeJSON).Route("/properties", func(r chi.Router) {
		r.Get("/", properties(app))
	})

	return r
}

func index(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, `Veds

* /namespaces`); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func namespaces(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		namespaces, err := service.Namespaces(r.Context(), app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(namespaces); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func kinds(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys, err := service.Kinds(r.Context(), app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(keys); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func properties(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys, err := service.Properties(r.Context(), app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(keys); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
