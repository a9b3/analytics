package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/esayemm/analytics/db"
	"github.com/go-chi/chi"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// ErrAlreadyExist is returned if a resource being created already exists
	ErrAlreadyExist = errors.New("already exists")
)

type getResponse struct {
	Results []db.Application `json:"results"`
}

func mgoQFromURLQ(q url.Values) map[string]interface{} {
	mgoQuery := bson.M{}
	for k, _ := range q {
		if q.Get(k) != "" {
			mgoQuery[k] = q.Get(k)
		}
	}
	return mgoQuery
}

// Router returns a new restful router that can be mounted
func Router(col *mgo.Collection) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", createGet(col))
	r.Get("/{id}", createGetOne(col))
	r.Post("/", createPost(col))

	return r
}

func createGet(col *mgo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var docs []db.Application
		err := col.Find(mgoQFromURLQ(r.URL.Query())).All(&docs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		content, err := json.Marshal(getResponse{Results: docs})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(content)
	}
}

func createPost(col *mgo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doc := db.Application{}
		doc.UserID = r.Context().Value("userId").(string)

		err := json.NewDecoder(r.Body).Decode(&doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		size, err := col.Find(bson.M{"name": doc.Name}).Count()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if size != 0 {
			http.Error(w, ErrAlreadyExist.Error(), http.StatusBadRequest)
			return
		}

		err = col.Insert(doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		content, err := json.Marshal(doc)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(content)
	}
}

func createGetOne(col *mgo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var doc db.Application
		err := col.FindId(bson.ObjectIdHex(id)).One(&doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		content, err := json.Marshal(doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(content)
	}
}
