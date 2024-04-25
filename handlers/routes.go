package handlers

import (
    "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {

    r := chi.NewRouter()

    r.Use(middleware.Logger)

    r.Get("/", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("No fotos here ..."))
    })

    return r
}
