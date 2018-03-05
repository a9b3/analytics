package app

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/esayemm/analytics/database"
	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2/bson"
)

func Router(applicationStore *database.ApplicationStore) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", createGet(applicationStore))
	r.Get("/{id}", createGetOne(applicationStore))
	r.Post("/", createPost(applicationStore))

	return r
}

type getPayload struct {
	Results []database.Application `json:"results"`
}

func createGet(applicationStore *database.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results, err := applicationStore.Get(mgoQueryFromUrlQuery(r.URL.Query()))
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

func mgoQueryFromUrlQuery(q url.Values) map[string]interface{} {
	mgoQuery := bson.M{}
	for k, _ := range q {
		if q.Get(k) != "" {
			mgoQuery[k] = q.Get(k)
		}
	}
	return mgoQuery
}

func createPost(applicationStore *database.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createdApp := database.Application{}
		createdApp.UserID = r.Context().Value("userId").(string)
		if err := json.NewDecoder(r.Body).Decode(&createdApp); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		results, err := applicationStore.Get(bson.M{"name": createdApp.Name})
		if err != nil || len(results) > 0 {
			http.Error(w, `Given name "`+createdApp.Name+`" already exists`, http.StatusInternalServerError)
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

func createGetOne(applicationStore *database.ApplicationStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		result, err := applicationStore.Get(bson.M{"_id": bson.ObjectIdHex(id)})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		content, err := json.Marshal(result[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(content)
	}
}
