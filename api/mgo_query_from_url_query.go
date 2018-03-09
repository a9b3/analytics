package api

import (
	"net/url"

	"gopkg.in/mgo.v2/bson"
)

func MgoQueryFromUrlQuery(q url.Values) map[string]interface{} {
	mgoQuery := bson.M{}
	for k, _ := range q {
		if q.Get(k) != "" {
			mgoQuery[k] = q.Get(k)
		}
	}
	return mgoQuery
}
