package middleware

import (
	"net/http"

	"github.com/esayemm/analytics/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func LazyCreateUser(col *mgo.Collection) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userId").(string)

			if userId == "" {
				http.Error(w, "userId must be provided", http.StatusInternalServerError)
				return
			}

			var users []db.User
			err := col.Find(bson.M{"_id": userId}).All(&users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if len(users) == 0 {
				err := col.Insert(&db.User{ID: userId})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
