package handler

import (
	"encoding/json"
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
	r.With(vmiddleware.ContentTypeJSON).Get("/entities/{kind}", entities(app))
	r.Get("/*", index(app))

	return r
}

// index returns the static file handler
func index(app *application.App) http.HandlerFunc {
	root := http.Dir(app.Config.Public)
	fs := http.FileServer(root)
	return func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
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
