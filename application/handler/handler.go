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
	r.With(vmiddleware.ContentTypeJSON).Get("/entities/{kind}", entities(app))

	return r
}

func index(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, `Veds

* /{kind}`); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func entities(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kind := chi.URLParam(r, "kind")
		entities, err := service.Entities(r.Context(), app, kind)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(entities); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
