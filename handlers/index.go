package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/eupedrosa/Fotos/pkg/tmpl"
)

func IndexRoutes(r chi.Router) {

    r.Get("/", func(w http.ResponseWriter, r *http.Request){

        t := tmpl.TmplHandle()
        t.Execute(w, "index", "")

    })

}
