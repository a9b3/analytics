package app

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// err := json.NewDecoder(r.Body).Decode(&u)
		// if err != nil {
		// 	http.Error()
		// 	return
		// }

		w.Write([]byte("hello"))
	})

	return r
}
