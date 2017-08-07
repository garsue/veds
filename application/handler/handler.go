package handler

import (
	"fmt"
	"net/http"

	"github.com/garsue/veds/application"
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
	r.HandleFunc("/", index(app))

	return r
}

func index(app *application.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		namespaces, err := service.Namespaces(r.Context(), app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Namespaces %v", namespaces)
	}
}
