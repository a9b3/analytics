package app

import (
	"encoding/json"
	"net/http"

	"github.com/esayemm/analytics/database"
	"github.com/go-chi/chi"
)

func Router(applicationStore *database.ApplicationStore) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", createGet(applicationStore))
	r.Post("/", createPost(applicationStore))

	return r
}

type getPayload struct {
	Results []database.Application `json:"results"`
}

func createGet(applicationStore *database.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		results, err := applicationStore.Get(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		content, err := json.Marshal(getPayload{results})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(content)
	}
}

func createPost(applicationStore *database.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createdApp := database.Application{}
		if err := json.NewDecoder(r.Body).Decode(&createdApp); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := applicationStore.Create(&createdApp); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		content, err := json.Marshal(createdApp)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(content)
	}
}
