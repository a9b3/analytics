package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

// TODO implement mongo operators
// eg. age=gte:10 => { age: { $gte: 10 } }
//     age=gte:10,lt:20 => { age: { $gte: 10, $lt: 20 } }
func valueToBson(s string) interface{} {
	operators := strings.Split(s, ",")
	if len(operators) == 1 {
		fmt.Println("here")
		return s
	}
	return s
}

func mgoQFromURLQ(q url.Values) map[string]interface{} {
	mgoQuery := bson.M{}
	for k, _ := range q {
		switch k {
		case "sort", "limit":
			continue
		}
		mgoQuery[k] = valueToBson(q.Get(k))
	}
	return mgoQuery
}

func buildQuery(c *mgo.Collection, q url.Values) *mgo.Query {
	query := c.Find(mgoQFromURLQ(q))

	sortBys := strings.Split(q.Get("sort"), ",")
	if len(sortBys) > 0 {
		query = query.Sort(sortBys...)
	}

	limit := q.Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err == nil {
		query = query.Limit(limitInt)
	}

	return query
}

// Router returns a new restful router that can be mounted
func Router(col *mgo.Collection) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", createGet(col))
	r.Get("/{id}", createGetByID(col))
	r.Post("/", createPost(col))

	return r
}

func createGet(col *mgo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		docs := make([]db.Application, 0)

		q := buildQuery(col, r.URL.Query())
		err := q.All(&docs)
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
		doc := db.Application{ID: bson.NewObjectId()}
		doc.UserID = r.Context().Value("userId").(string)

		err := json.NewDecoder(r.Body).Decode(&doc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = col.Insert(doc)
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

func createGetByID(col *mgo.Collection) http.HandlerFunc {
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
