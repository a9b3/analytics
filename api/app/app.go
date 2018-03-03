package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	return r
}
