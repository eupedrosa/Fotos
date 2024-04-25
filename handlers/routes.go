package handlers

import (
    "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {

    r := chi.NewRouter()

    r.Use(middleware.Logger)

    // Route for static data under /static
    fserver := http.FileServer(http.Dir("assets/static"))
    r.Handle("/static/*", http.StripPrefix("/static/", fserver))

    r.Route("/", IndexRoutes)

    return r
}
