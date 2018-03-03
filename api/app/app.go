package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	return r
}
